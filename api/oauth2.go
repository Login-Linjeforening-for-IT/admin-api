package api

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"crypto/tls"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"gitlab.login.no/tekkom/web/beehive/admin-api/sessionstore"
	"golang.org/x/oauth2"
)

type userInfo struct {
	ID    string   `json:"id"`
	Name  string   `json:"name"`
	Email string   `json:"email"`
	Roles []string `json:"roles"`
	// Picture string `json:"picture"`
}

type oauth2Config struct {
	oauth2.Config
	UserInfoEndpoint string
	RevokeEndpoint   string
	stateExpiration  time.Duration
	provider         string
}

var stateExpiration = 20 * time.Minute

func (conf *oauth2Config) generateStateOauthCookie(ctx *gin.Context) string {
	state := generateOauthState()
	ctx.SetCookie("oauthstate", state, int(conf.stateExpiration.Seconds()), "/", "", false, true)
	return state
}

func generateOauthState() string {
	return xid.New().String()
}

// Custom HTTP client with TLS verification disabled
func selfSignedClient() *http.Client {
    return &http.Client{
        Transport: &http.Transport{
            TLSClientConfig: &tls.Config{
                InsecureSkipVerify: true,
            },
        },
    }
}

// exchange code for token
func (conf oauth2Config) getToken(code string) (*oauth2.Token, error) {
    // Create a context with a custom HTTP client
    insecureClient := selfSignedClient()
    ctx := context.WithValue(context.Background(), oauth2.HTTPClient, insecureClient)

    // Exchange the code for a token
    token, err := conf.Exchange(ctx, code)
    if err != nil {
        return nil, fmt.Errorf("code exchange wrong: %w", err)
    }

    return token, nil
}

func (server *Server) oauth2Login(ctx *gin.Context) {
	oauthState := server.oauth2Config.generateStateOauthCookie(ctx)
	u := server.oauth2Config.AuthCodeURL(oauthState)
	ctx.Redirect(http.StatusTemporaryRedirect, u)
}

type getUserInfoFunc func(ctx context.Context, token *oauth2.Token) (userInfo, error)

type oauth2FallbackResponse struct {
	AccessToken           string                     `json:"access_token"`
	AccessTokenExpiresAt  time.Time                  `json:"access_token_expires_at"`
	RefreshToken          string                     `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time                  `json:"refresh_token_expires_at"`
	User                  oauth2FallbackResponseUser `json:"user"`
}

type oauth2FallbackResponseUser struct {
	ID    string   `json:"id"`
	Name  string   `json:"name"`
	Roles []string `json:"roles"`
}

func (server *Server) oauth2Fallback(provider string, getUserInfo getUserInfoFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		oauthState, err := ctx.Request.Cookie("oauthstate")
		if err != nil || ctx.Request.FormValue("state") != oauthState.Value {
			err = fmt.Errorf("error - " + ctx.Request.FormValue("state"))

			server.writeError(ctx, http.StatusUnauthorized, fmt.Errorf("oauth2Fallback, Cookie or FormValue - %w", err))
			return
		}

		code := ctx.Request.FormValue("code")
		userToken, err := server.oauth2Config.getToken(code)
		if err != nil {
			server.writeError(ctx, http.StatusInternalServerError, fmt.Errorf("oauth2Fallback, getToken  - %w", err))
			return
		}

		userInfo, err := getUserInfo(ctx, userToken)
		if err != nil {
			server.writeError(ctx, http.StatusInternalServerError, fmt.Errorf("oauth2Fallback, getUserInfo  - %w", err))
			return
		}

		// TODO(session-store): Consider upserting user
		// user, err := server.service.UpsertUser(ctx, db.UpsertUserParams{
		// 	Provider:       provider,
		// 	ProviderUserID: userInfo.ID,
		// 	Name:           userInfo.Name,
		// 	Email:          userInfo.Email,
		// })
		// if err != nil {
		// 	server.writeDBError(ctx, err)
		// 	return
		// }

		accessToken, accessTokenPayload, err := server.createAccessToken(ctx, userInfo.ID, userInfo.Roles)
		if err != nil {
			server.writeError(ctx, http.StatusInternalServerError, fmt.Errorf("oauth2Fallback, createAccessToken  - %w", err))
			return
		}

		refreshToken, refreshTokenPayload, err := server.createRefreshToken(ctx, userInfo.ID, userInfo.Roles)
		if err != nil {
			server.writeError(ctx, http.StatusInternalServerError, fmt.Errorf("oauth2Fallback, createRefreshToken  - %w", err))
			return
		}

		err = server.sessionstore.CreateSession(ctx, sessionstore.CreateSessionParams{
			ID:           refreshTokenPayload.ID,
			UID:          userInfo.ID,
			RefreshToken: refreshToken,
			UserAgent:    ctx.Request.UserAgent(),
			ClientIP:     ctx.ClientIP(),
			ExpiresAt:    refreshTokenPayload.ExpiresAt,
		})
		if err != nil {
			server.writeDBError(ctx, err)
			return
		}

		setAccessTokenCookie(ctx, accessToken, accessTokenPayload)
		setRefreshTokenCookie(ctx, refreshToken, refreshTokenPayload)

		ctx.JSON(http.StatusOK, oauth2FallbackResponse{
			AccessToken:           accessToken,
			AccessTokenExpiresAt:  accessTokenPayload.ExpiresAt,
			RefreshToken:          refreshToken,
			RefreshTokenExpiresAt: refreshTokenPayload.ExpiresAt,
			User: oauth2FallbackResponseUser{
				ID:    userInfo.ID,
				Name:  userInfo.Name,
				Roles: accessTokenPayload.Roles,
			},
		})
	}
}