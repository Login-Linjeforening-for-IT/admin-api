// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: events.sql

package db

import (
	"context"
	"time"

	"github.com/guregu/null/zero"
	"github.com/lib/pq"
)

const addAudienceToEvent = `-- name: AddAudienceToEvent :exec
INSERT INTO "event_audience_relation" ("event", "audience") VALUES ($1, $2) ON CONFLICT DO NOTHING
`

type AddAudienceToEventParams struct {
	Event    int32 `json:"event"`
	Audience int32 `json:"audience"`
}

func (q *Queries) AddAudienceToEvent(ctx context.Context, arg AddAudienceToEventParams) error {
	_, err := q.db.ExecContext(ctx, addAudienceToEvent, arg.Event, arg.Audience)
	return err
}

const addOrganizationToEvent = `-- name: AddOrganizationToEvent :exec
INSERT INTO "event_organization_relation" ("event", "organization") VALUES ($1, $2) ON CONFLICT DO NOTHING
`

type AddOrganizationToEventParams struct {
	Event        int32  `json:"event"`
	Organization string `json:"organization"`
}

func (q *Queries) AddOrganizationToEvent(ctx context.Context, arg AddOrganizationToEventParams) error {
	_, err := q.db.ExecContext(ctx, addOrganizationToEvent, arg.Event, arg.Organization)
	return err
}

const createEvent = `-- name: CreateEvent :one

INSERT INTO "event" (
    "visible",
    "name_no", "name_en",
    "description_no", "description_en",
    "informational_no", "informational_en",
    "time_type", "time_start", "time_end", "time_publish",
    "time_signup_release", "time_signup_deadline",
    "canceled", "digital", "highlight",
    "image_small", "image_banner",
    "link_facebook", "link_discord", "link_signup", "link_stream",
    "capacity", "full",
    "category", "location", "parent", "rule"
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28)
RETURNING id, visible, name_no, name_en, description_no, description_en, informational_no, informational_en, time_type, time_start, time_end, time_publish, time_signup_release, time_signup_deadline, canceled, digital, highlight, image_small, image_banner, link_facebook, link_discord, link_signup, link_stream, capacity, "full", category, location, parent, rule, updated_at, created_at, deleted_at
`

type CreateEventParams struct {
	Visible            bool         `json:"visible"`
	NameNo             string       `json:"name_no"`
	NameEn             zero.String  `json:"name_en"`
	DescriptionNo      string       `json:"description_no"`
	DescriptionEn      zero.String  `json:"description_en"`
	InformationalNo    zero.String  `json:"informational_no"`
	InformationalEn    zero.String  `json:"informational_en"`
	TimeType           TimeTypeEnum `json:"time_type"`
	TimeStart          time.Time    `json:"time_start"`
	TimeEnd            zero.Time    `json:"time_end"`
	TimePublish        zero.Time    `json:"time_publish"`
	TimeSignupRelease  zero.Time    `json:"time_signup_release"`
	TimeSignupDeadline zero.Time    `json:"time_signup_deadline"`
	Canceled           bool         `json:"canceled"`
	Digital            bool         `json:"digital"`
	Highlight          bool         `json:"highlight"`
	ImageSmall         string       `json:"image_small"`
	ImageBanner        string       `json:"image_banner"`
	LinkFacebook       zero.String  `json:"link_facebook"`
	LinkDiscord        zero.String  `json:"link_discord"`
	LinkSignup         zero.String  `json:"link_signup"`
	LinkStream         zero.String  `json:"link_stream"`
	Capacity           zero.Int     `json:"capacity"`
	Full               bool         `json:"full"`
	Category           int32        `json:"category"`
	Location           zero.Int     `json:"location"`
	Parent             zero.Int     `json:"parent"`
	Rule               zero.Int     `json:"rule"`
}

// sqlc.embed(event)
func (q *Queries) CreateEvent(ctx context.Context, arg CreateEventParams) (Event, error) {
	row := q.db.QueryRowContext(ctx, createEvent,
		arg.Visible,
		arg.NameNo,
		arg.NameEn,
		arg.DescriptionNo,
		arg.DescriptionEn,
		arg.InformationalNo,
		arg.InformationalEn,
		arg.TimeType,
		arg.TimeStart,
		arg.TimeEnd,
		arg.TimePublish,
		arg.TimeSignupRelease,
		arg.TimeSignupDeadline,
		arg.Canceled,
		arg.Digital,
		arg.Highlight,
		arg.ImageSmall,
		arg.ImageBanner,
		arg.LinkFacebook,
		arg.LinkDiscord,
		arg.LinkSignup,
		arg.LinkStream,
		arg.Capacity,
		arg.Full,
		arg.Category,
		arg.Location,
		arg.Parent,
		arg.Rule,
	)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.Visible,
		&i.NameNo,
		&i.NameEn,
		&i.DescriptionNo,
		&i.DescriptionEn,
		&i.InformationalNo,
		&i.InformationalEn,
		&i.TimeType,
		&i.TimeStart,
		&i.TimeEnd,
		&i.TimePublish,
		&i.TimeSignupRelease,
		&i.TimeSignupDeadline,
		&i.Canceled,
		&i.Digital,
		&i.Highlight,
		&i.ImageSmall,
		&i.ImageBanner,
		&i.LinkFacebook,
		&i.LinkDiscord,
		&i.LinkSignup,
		&i.LinkStream,
		&i.Capacity,
		&i.Full,
		&i.Category,
		&i.Location,
		&i.Parent,
		&i.Rule,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getEvent = `-- name: GetEvent :one
SELECT id, visible, name_no, name_en, description_no, description_en, informational_no, informational_en, time_type, time_start, time_end, time_publish, time_signup_release, time_signup_deadline, canceled, digital, highlight, image_small, image_banner, link_facebook, link_discord, link_signup, link_stream, capacity, "full", category, location, parent, rule, updated_at, created_at, deleted_at FROM "event" WHERE "id" = $1::int LIMIT 1
`

func (q *Queries) GetEvent(ctx context.Context, id int32) (Event, error) {
	row := q.db.QueryRowContext(ctx, getEvent, id)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.Visible,
		&i.NameNo,
		&i.NameEn,
		&i.DescriptionNo,
		&i.DescriptionEn,
		&i.InformationalNo,
		&i.InformationalEn,
		&i.TimeType,
		&i.TimeStart,
		&i.TimeEnd,
		&i.TimePublish,
		&i.TimeSignupRelease,
		&i.TimeSignupDeadline,
		&i.Canceled,
		&i.Digital,
		&i.Highlight,
		&i.ImageSmall,
		&i.ImageBanner,
		&i.LinkFacebook,
		&i.LinkDiscord,
		&i.LinkSignup,
		&i.LinkStream,
		&i.Capacity,
		&i.Full,
		&i.Category,
		&i.Location,
		&i.Parent,
		&i.Rule,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getEvents = `-- name: GetEvents :many
SELECT e."id", e."visible",
        e."name_no", e."name_en",
        e."time_type", e."time_start", e."time_end", e."time_publish",
        e."canceled", e."link_signup", e."capacity", e."full",
        c."name_no" AS category_name_no, c."name_en"  AS category_name_en, 
        l."name_no" AS location_name_no, l."name_en" AS location_name_en,
        e."updated_at", (e."deleted_at" IS NOT NULL)::bool AS is_deleted,
        array(
            SELECT COALESCE(NULLIF("name_en", ''), "name_no") FROM "audience"
            INNER JOIN "event_audience_relation" AS ear ON ear."audience" = "audience"."id"
            WHERE ear."event" = e."id"
		)::varchar[] AS audiences,
		array(
			SELECT COALESCE(NULLIF("name_en", ''), "name_no") FROM "organization"
			INNER JOIN "event_organization_relation" AS eor ON eor."organization" = "organization"."shortname"
			WHERE eor."event" = e."id"
		)::varchar[] AS organizers
    FROM "event" AS e
    INNER JOIN "category" AS c ON e."category" = c."id"
    LEFT OUTER JOIN "location" AS l ON e."location" = l."id"
    WHERE ($1::bool OR ((e."time_end" IS NOT NULL AND e."time_end" > now()) OR (e."time_start" > now() - interval '1 day')))
    ORDER BY e."id"
    LIMIT $3::int
    OFFSET $2::int
`

type GetEventsParams struct {
	Historical bool  `json:"historical"`
	Offset     int32 `json:"offset"`
	Limit      int32 `json:"limit"`
}

type GetEventsRow struct {
	ID             int32        `json:"id"`
	Visible        bool         `json:"visible"`
	NameNo         string       `json:"name_no"`
	NameEn         zero.String  `json:"name_en"`
	TimeType       TimeTypeEnum `json:"time_type"`
	TimeStart      time.Time    `json:"time_start"`
	TimeEnd        zero.Time    `json:"time_end"`
	TimePublish    zero.Time    `json:"time_publish"`
	Canceled       bool         `json:"canceled"`
	LinkSignup     zero.String  `json:"link_signup"`
	Capacity       zero.Int     `json:"capacity"`
	Full           bool         `json:"full"`
	CategoryNameNo string       `json:"category_name_no"`
	CategoryNameEn zero.String  `json:"category_name_en"`
	LocationNameNo zero.String  `json:"location_name_no"`
	LocationNameEn zero.String  `json:"location_name_en"`
	UpdatedAt      time.Time    `json:"updated_at"`
	IsDeleted      bool         `json:"is_deleted"`
	Audiences      []string     `json:"audiences"`
	Organizers     []string     `json:"organizers"`
}

func (q *Queries) GetEvents(ctx context.Context, arg GetEventsParams) ([]GetEventsRow, error) {
	rows, err := q.db.QueryContext(ctx, getEvents, arg.Historical, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetEventsRow{}
	for rows.Next() {
		var i GetEventsRow
		if err := rows.Scan(
			&i.ID,
			&i.Visible,
			&i.NameNo,
			&i.NameEn,
			&i.TimeType,
			&i.TimeStart,
			&i.TimeEnd,
			&i.TimePublish,
			&i.Canceled,
			&i.LinkSignup,
			&i.Capacity,
			&i.Full,
			&i.CategoryNameNo,
			&i.CategoryNameEn,
			&i.LocationNameNo,
			&i.LocationNameEn,
			&i.UpdatedAt,
			&i.IsDeleted,
			pq.Array(&i.Audiences),
			pq.Array(&i.Organizers),
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const removeAudienceFromEvent = `-- name: RemoveAudienceFromEvent :exec
DELETE FROM "event_audience_relation" WHERE ("event", "audience") = ($1::int, $2::int)
`

type RemoveAudienceFromEventParams struct {
	Event    int32 `json:"event"`
	Audience int32 `json:"audience"`
}

func (q *Queries) RemoveAudienceFromEvent(ctx context.Context, arg RemoveAudienceFromEventParams) error {
	_, err := q.db.ExecContext(ctx, removeAudienceFromEvent, arg.Event, arg.Audience)
	return err
}

const removeOrganizationFromEvent = `-- name: RemoveOrganizationFromEvent :exec
DELETE FROM "event_organization_relation" WHERE ("event", "organization") = ($1::int, $2::varchar)
`

type RemoveOrganizationFromEventParams struct {
	Event        int32  `json:"event"`
	Organization string `json:"organization"`
}

func (q *Queries) RemoveOrganizationFromEvent(ctx context.Context, arg RemoveOrganizationFromEventParams) error {
	_, err := q.db.ExecContext(ctx, removeOrganizationFromEvent, arg.Event, arg.Organization)
	return err
}

const softDeleteEvent = `-- name: SoftDeleteEvent :one
UPDATE "event"
SET
    "deleted_at" = now(),
    "updated_at" = now()
WHERE "id" = $1::int
RETURNING id, visible, name_no, name_en, description_no, description_en, informational_no, informational_en, time_type, time_start, time_end, time_publish, time_signup_release, time_signup_deadline, canceled, digital, highlight, image_small, image_banner, link_facebook, link_discord, link_signup, link_stream, capacity, "full", category, location, parent, rule, updated_at, created_at, deleted_at
`

func (q *Queries) SoftDeleteEvent(ctx context.Context, id int32) (Event, error) {
	row := q.db.QueryRowContext(ctx, softDeleteEvent, id)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.Visible,
		&i.NameNo,
		&i.NameEn,
		&i.DescriptionNo,
		&i.DescriptionEn,
		&i.InformationalNo,
		&i.InformationalEn,
		&i.TimeType,
		&i.TimeStart,
		&i.TimeEnd,
		&i.TimePublish,
		&i.TimeSignupRelease,
		&i.TimeSignupDeadline,
		&i.Canceled,
		&i.Digital,
		&i.Highlight,
		&i.ImageSmall,
		&i.ImageBanner,
		&i.LinkFacebook,
		&i.LinkDiscord,
		&i.LinkSignup,
		&i.LinkStream,
		&i.Capacity,
		&i.Full,
		&i.Category,
		&i.Location,
		&i.Parent,
		&i.Rule,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const updateEvent = `-- name: UpdateEvent :one
UPDATE "event"
SET
    "visible" = COALESCE($1, visible),
    "name_no" = COALESCE($2, name_no),
    "name_en" = COALESCE($3, name_en),
    "description_no" = COALESCE($4, description_no),
    "description_en" = COALESCE($5, description_en),
    "informational_no" = COALESCE($6, informational_no),
    "informational_en" = COALESCE($7, informational_en),
    "time_type" = COALESCE($8, time_type),
    "time_start" = COALESCE($9, time_start),
    "time_end" = COALESCE($10, time_end),
    "time_publish" = COALESCE($11, time_publish),
    "time_signup_release" = COALESCE($12, time_signup_release),
    "time_signup_deadline" = COALESCE($13, time_signup_deadline),
    "canceled" = COALESCE($14, canceled),
    "digital" = COALESCE($15, digital),
    "highlight" = COALESCE($16, highlight),
    "image_small" = COALESCE($17, image_small),
    "image_banner" = COALESCE($18, image_banner),
    "link_facebook" = COALESCE($19, link_facebook),
    "link_discord" = COALESCE($20, link_discord),
    "link_signup" = COALESCE($21, link_signup),
    "link_stream" = COALESCE($22, link_stream),
    "capacity" = COALESCE($23, capacity),
    "full" = COALESCE($24, "full"),
    "category" = COALESCE($25, category),
    "location" = COALESCE($26, "location"),
    "parent" = COALESCE($27, parent),
    "rule" = COALESCE($28, rule),
    "updated_at" = now()
WHERE "id" = $29::int
RETURNING id, visible, name_no, name_en, description_no, description_en, informational_no, informational_en, time_type, time_start, time_end, time_publish, time_signup_release, time_signup_deadline, canceled, digital, highlight, image_small, image_banner, link_facebook, link_discord, link_signup, link_stream, capacity, "full", category, location, parent, rule, updated_at, created_at, deleted_at
`

type UpdateEventParams struct {
	Visible            zero.Bool        `json:"visible"`
	NameNo             zero.String      `json:"name_no"`
	NameEn             zero.String      `json:"name_en"`
	DescriptionNo      zero.String      `json:"description_no"`
	DescriptionEn      zero.String      `json:"description_en"`
	InformationalNo    zero.String      `json:"informational_no"`
	InformationalEn    zero.String      `json:"informational_en"`
	TimeType           NullTimeTypeEnum `json:"time_type"`
	TimeStart          zero.Time        `json:"time_start"`
	TimeEnd            zero.Time        `json:"time_end"`
	TimePublish        zero.Time        `json:"time_publish"`
	TimeSignupRelease  zero.Time        `json:"time_signup_release"`
	TimeSignupDeadline zero.Time        `json:"time_signup_deadline"`
	Canceled           zero.Bool        `json:"canceled"`
	Digital            zero.Bool        `json:"digital"`
	Highlight          zero.Bool        `json:"highlight"`
	ImageSmall         zero.String      `json:"image_small"`
	ImageBanner        zero.String      `json:"image_banner"`
	LinkFacebook       zero.String      `json:"link_facebook"`
	LinkDiscord        zero.String      `json:"link_discord"`
	LinkSignup         zero.String      `json:"link_signup"`
	LinkStream         zero.String      `json:"link_stream"`
	Capacity           zero.Int         `json:"capacity"`
	Full               zero.Bool        `json:"full"`
	Category           zero.Int         `json:"category"`
	Location           zero.Int         `json:"location"`
	Parent             zero.Int         `json:"parent"`
	Rule               zero.Int         `json:"rule"`
	ID                 int32            `json:"id"`
}

func (q *Queries) UpdateEvent(ctx context.Context, arg UpdateEventParams) (Event, error) {
	row := q.db.QueryRowContext(ctx, updateEvent,
		arg.Visible,
		arg.NameNo,
		arg.NameEn,
		arg.DescriptionNo,
		arg.DescriptionEn,
		arg.InformationalNo,
		arg.InformationalEn,
		arg.TimeType,
		arg.TimeStart,
		arg.TimeEnd,
		arg.TimePublish,
		arg.TimeSignupRelease,
		arg.TimeSignupDeadline,
		arg.Canceled,
		arg.Digital,
		arg.Highlight,
		arg.ImageSmall,
		arg.ImageBanner,
		arg.LinkFacebook,
		arg.LinkDiscord,
		arg.LinkSignup,
		arg.LinkStream,
		arg.Capacity,
		arg.Full,
		arg.Category,
		arg.Location,
		arg.Parent,
		arg.Rule,
		arg.ID,
	)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.Visible,
		&i.NameNo,
		&i.NameEn,
		&i.DescriptionNo,
		&i.DescriptionEn,
		&i.InformationalNo,
		&i.InformationalEn,
		&i.TimeType,
		&i.TimeStart,
		&i.TimeEnd,
		&i.TimePublish,
		&i.TimeSignupRelease,
		&i.TimeSignupDeadline,
		&i.Canceled,
		&i.Digital,
		&i.Highlight,
		&i.ImageSmall,
		&i.ImageBanner,
		&i.LinkFacebook,
		&i.LinkDiscord,
		&i.LinkSignup,
		&i.LinkStream,
		&i.Capacity,
		&i.Full,
		&i.Category,
		&i.Location,
		&i.Parent,
		&i.Rule,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}
