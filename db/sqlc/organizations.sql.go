// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: organizations.sql

package db

import (
	"context"
	"time"

	"github.com/guregu/null/zero"
)

const createOrganization = `-- name: CreateOrganization :one
INSERT INTO "organization" (
    "shortname",
    "name_no", "name_en",
    "description_no", "description_en",
    "link_homepage", "link_linkedin", "link_facebook", "link_instagram",
    "logo"
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING shortname, name_no, name_en, description_no, description_en, link_homepage, link_linkedin, link_facebook, link_instagram, logo, updated_at, created_at, deleted_at
`

type CreateOrganizationParams struct {
	Shortname     string      `json:"shortname"`
	NameNo        string      `json:"name_no"`
	NameEn        zero.String `json:"name_en"`
	DescriptionNo string      `json:"description_no"`
	DescriptionEn zero.String `json:"description_en"`
	LinkHomepage  zero.String `json:"link_homepage"`
	LinkLinkedin  zero.String `json:"link_linkedin"`
	LinkFacebook  zero.String `json:"link_facebook"`
	LinkInstagram zero.String `json:"link_instagram"`
	Logo          zero.String `json:"logo"`
}

func (q *Queries) CreateOrganization(ctx context.Context, arg CreateOrganizationParams) (Organization, error) {
	row := q.db.QueryRowContext(ctx, createOrganization,
		arg.Shortname,
		arg.NameNo,
		arg.NameEn,
		arg.DescriptionNo,
		arg.DescriptionEn,
		arg.LinkHomepage,
		arg.LinkLinkedin,
		arg.LinkFacebook,
		arg.LinkInstagram,
		arg.Logo,
	)
	var i Organization
	err := row.Scan(
		&i.Shortname,
		&i.NameNo,
		&i.NameEn,
		&i.DescriptionNo,
		&i.DescriptionEn,
		&i.LinkHomepage,
		&i.LinkLinkedin,
		&i.LinkFacebook,
		&i.LinkInstagram,
		&i.Logo,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getOrganization = `-- name: GetOrganization :one
SELECT shortname, name_no, name_en, description_no, description_en, link_homepage, link_linkedin, link_facebook, link_instagram, logo, updated_at, created_at, deleted_at FROM "organization" WHERE "shortname" = $1::text LIMIT 1
`

func (q *Queries) GetOrganization(ctx context.Context, shortname string) (Organization, error) {
	row := q.db.QueryRowContext(ctx, getOrganization, shortname)
	var i Organization
	err := row.Scan(
		&i.Shortname,
		&i.NameNo,
		&i.NameEn,
		&i.DescriptionNo,
		&i.DescriptionEn,
		&i.LinkHomepage,
		&i.LinkLinkedin,
		&i.LinkFacebook,
		&i.LinkInstagram,
		&i.Logo,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getOrganizations = `-- name: GetOrganizations :many
SELECT "shortname", "name_no", "name_en", "link_homepage", "logo", "updated_at", "deleted_at"
FROM "organization"
ORDER BY "shortname"
LIMIT $2::int
OFFSET $1::int
`

type GetOrganizationsParams struct {
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
}

type GetOrganizationsRow struct {
	Shortname    string      `json:"shortname"`
	NameNo       string      `json:"name_no"`
	NameEn       zero.String `json:"name_en"`
	LinkHomepage zero.String `json:"link_homepage"`
	Logo         zero.String `json:"logo"`
	UpdatedAt    time.Time   `json:"updated_at"`
	DeletedAt    zero.Time   `json:"deleted_at"`
}

func (q *Queries) GetOrganizations(ctx context.Context, arg GetOrganizationsParams) ([]GetOrganizationsRow, error) {
	rows, err := q.db.QueryContext(ctx, getOrganizations, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetOrganizationsRow{}
	for rows.Next() {
		var i GetOrganizationsRow
		if err := rows.Scan(
			&i.Shortname,
			&i.NameNo,
			&i.NameEn,
			&i.LinkHomepage,
			&i.Logo,
			&i.UpdatedAt,
			&i.DeletedAt,
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

const getOrganizationsOfEvent = `-- name: GetOrganizationsOfEvent :many
SELECT org.shortname, org.name_no, org.name_en, org.description_no, org.description_en, org.link_homepage, org.link_linkedin, org.link_facebook, org.link_instagram, org.logo, org.updated_at, org.created_at, org.deleted_at FROM "event_organization_relation"
    INNER JOIN "organization" AS org ON "event_organization_relation"."organization" = org."shortname"
    WHERE "event_organization_relation"."event" = $1::int
`

func (q *Queries) GetOrganizationsOfEvent(ctx context.Context, eventID int32) ([]Organization, error) {
	rows, err := q.db.QueryContext(ctx, getOrganizationsOfEvent, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Organization{}
	for rows.Next() {
		var i Organization
		if err := rows.Scan(
			&i.Shortname,
			&i.NameNo,
			&i.NameEn,
			&i.DescriptionNo,
			&i.DescriptionEn,
			&i.LinkHomepage,
			&i.LinkLinkedin,
			&i.LinkFacebook,
			&i.LinkInstagram,
			&i.Logo,
			&i.UpdatedAt,
			&i.CreatedAt,
			&i.DeletedAt,
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

const updateOrganization = `-- name: UpdateOrganization :one
UPDATE "organization"
SET
    "name_no" = COALESCE($1, name_no),
    "name_en" = COALESCE($2, name_en),
    "description_no" = COALESCE($3, description_no),
    "description_en" = COALESCE($4, description_en),
    "link_homepage" = COALESCE($5, link_homepage),
    "link_linkedin" = COALESCE($6, link_linkedin),
    "link_facebook" = COALESCE($7, link_facebook),
    "link_instagram" = COALESCE($8, link_instagram),
    "logo" = COALESCE($9, logo),
    "updated_at" = now()
WHERE "shortname" = $10::text RETURNING shortname, name_no, name_en, description_no, description_en, link_homepage, link_linkedin, link_facebook, link_instagram, logo, updated_at, created_at, deleted_at
`

type UpdateOrganizationParams struct {
	NameNo        zero.String `json:"name_no"`
	NameEn        zero.String `json:"name_en"`
	DescriptionNo zero.String `json:"description_no"`
	DescriptionEn zero.String `json:"description_en"`
	LinkHomepage  zero.String `json:"link_homepage"`
	LinkLinkedin  zero.String `json:"link_linkedin"`
	LinkFacebook  zero.String `json:"link_facebook"`
	LinkInstagram zero.String `json:"link_instagram"`
	Logo          zero.String `json:"logo"`
	Shortname     string      `json:"shortname"`
}

func (q *Queries) UpdateOrganization(ctx context.Context, arg UpdateOrganizationParams) (Organization, error) {
	row := q.db.QueryRowContext(ctx, updateOrganization,
		arg.NameNo,
		arg.NameEn,
		arg.DescriptionNo,
		arg.DescriptionEn,
		arg.LinkHomepage,
		arg.LinkLinkedin,
		arg.LinkFacebook,
		arg.LinkInstagram,
		arg.Logo,
		arg.Shortname,
	)
	var i Organization
	err := row.Scan(
		&i.Shortname,
		&i.NameNo,
		&i.NameEn,
		&i.DescriptionNo,
		&i.DescriptionEn,
		&i.LinkHomepage,
		&i.LinkLinkedin,
		&i.LinkFacebook,
		&i.LinkInstagram,
		&i.Logo,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}