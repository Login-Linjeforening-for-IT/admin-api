package api

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type errorResponse struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
	Type   string `json:"type"`
}

func newErrorResponse(status int, err error) errorResponse {
	res := errorResponse{
		Status: status,
	}

	if err != nil {
		res.Error = err.Error()
		res.Type = reflect.TypeOf(err).String()
	}

	return res
}

type RedactedError struct {
	ID      string `json:"id"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e *RedactedError) Error() string {
	return fmt.Sprintf("(id=%s) %s", e.ID, e.Message)
}

type NotFoundError struct {
	Message string `json:"message"`
}

func (e *NotFoundError) Error() string {
	return e.Message
}

type ValidationErrors struct {
	Errors []ValidationError `json:"errors"`
}

func (e *ValidationErrors) Error() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Found %d validation errors: ", len(e.Errors)))

	for i, err := range e.Errors {
		if i != 0 {
			sb.WriteString(", ")
		}

		sb.WriteString(err.Error())
	}

	return sb.String()
}

type ValidationError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation for field %s failed on tag %s", e.Field, e.Tag)
}

func (server *Server) writeError(ctx *gin.Context, status int, err error) {
	if status >= 500 {
		err = server.redactError(err)
	} else if status == http.StatusNotFound {
		err = &NotFoundError{
			Message: ctx.Request.URL.Path + " not found",
		}
	} else {
		switch uErr := err.(type) {
		case validator.ValidationErrors:
			errs := make([]ValidationError, 0, len(uErr))
			for _, fieldErr := range uErr {
				errs = append(errs, ValidationError{
					Field: fieldErr.Field(),
					Tag:   fieldErr.Tag(),
				})
			}
			err = &ValidationErrors{errs}
		}
	}

	ctx.JSON(status, newErrorResponse(status, err))
}

// If error is not a validation error, this functions as an alias for server.writeError(ctx, http.StatusBadRequest, err)
// However, if error is a validation error this function will attempt to replace the go field name with the json tag if it exists
func writeValidationError[T any](server *Server, ctx *gin.Context, err error) {
	switch uErr := err.(type) {
	case validator.ValidationErrors:
		var dummy T
		t := reflect.TypeOf(dummy)
		for t.Kind() == reflect.Ptr {
			t = t.Elem()
		}

		if t.Kind() != reflect.Struct {
			break
		}

		errs := make([]ValidationError, 0, len(uErr))
		for _, fieldErr := range uErr {
			field := fieldErr.Field()

			sf, ok := t.FieldByName(field)
			if ok {
				tagField, ok := sf.Tag.Lookup("json")
				if ok {
					field = tagField
				}
			}

			errs = append(errs, ValidationError{
				Field: field,
				Tag:   fieldErr.Tag(),
			})
		}
		err = &ValidationErrors{errs}
	}

	server.writeError(ctx, http.StatusBadRequest, err)
}

func (server *Server) redactError(err error) error {
	id := uuid.NewString()

	errChain := []string{}
	for chainErr := err; chainErr != nil; chainErr = errors.Unwrap(chainErr) {
		errChain = append(errChain, chainErr.Error())
	}

	server.logger.Error().Err(err).Str("error-id", id).Str("error-chain", strings.Join(errChain, " -- ")).Int("error-chain-length", len(errChain)).Send()

	return &RedactedError{
		ID:      id,
		Status:  http.StatusInternalServerError,
		Message: "Something went wrong. Contact admin if problem persists.",
	}
}