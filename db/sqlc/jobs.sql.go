// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: jobs.sql

package db

import (
	"context"
	"time"

	"github.com/guregu/null/zero"
)

const createJob = `-- name: CreateJob :one
INSERT INTO "job_advertisement" (
    "visible",
    "title_no", "title_en", 
    "position_title_no", "position_title_en", 
    "description_short_no", "description_short_en", 
    "description_long_no", "description_long_en", 
    "job_type", "time_publish", "application_deadline", "banner_image", 
    "organization", "application_url"
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) RETURNING id, visible, title_no, title_en, position_title_no, position_title_en, description_short_no, description_short_en, description_long_no, description_long_en, job_type, time_publish, application_deadline, banner_image, organization, application_url, updated_at, created_at, deleted_at
`

type CreateJobParams struct {
	Visible             bool        `json:"visible"`
	TitleNo             string      `json:"title_no"`
	TitleEn             zero.String `json:"title_en"`
	PositionTitleNo     string      `json:"position_title_no"`
	PositionTitleEn     zero.String `json:"position_title_en"`
	DescriptionShortNo  string      `json:"description_short_no"`
	DescriptionShortEn  zero.String `json:"description_short_en"`
	DescriptionLongNo   string      `json:"description_long_no"`
	DescriptionLongEn   zero.String `json:"description_long_en"`
	JobType             JobType     `json:"job_type"`
	TimePublish         time.Time   `json:"time_publish"`
	ApplicationDeadline time.Time   `json:"application_deadline"`
	BannerImage         zero.String `json:"banner_image"`
	Organization        string      `json:"organization"`
	ApplicationUrl      string      `json:"application_url"`
}

func (q *Queries) CreateJob(ctx context.Context, arg CreateJobParams) (JobAdvertisement, error) {
	row := q.db.QueryRowContext(ctx, createJob,
		arg.Visible,
		arg.TitleNo,
		arg.TitleEn,
		arg.PositionTitleNo,
		arg.PositionTitleEn,
		arg.DescriptionShortNo,
		arg.DescriptionShortEn,
		arg.DescriptionLongNo,
		arg.DescriptionLongEn,
		arg.JobType,
		arg.TimePublish,
		arg.ApplicationDeadline,
		arg.BannerImage,
		arg.Organization,
		arg.ApplicationUrl,
	)
	var i JobAdvertisement
	err := row.Scan(
		&i.ID,
		&i.Visible,
		&i.TitleNo,
		&i.TitleEn,
		&i.PositionTitleNo,
		&i.PositionTitleEn,
		&i.DescriptionShortNo,
		&i.DescriptionShortEn,
		&i.DescriptionLongNo,
		&i.DescriptionLongEn,
		&i.JobType,
		&i.TimePublish,
		&i.ApplicationDeadline,
		&i.BannerImage,
		&i.Organization,
		&i.ApplicationUrl,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getJob = `-- name: GetJob :one
SELECT job.id, job.visible, job.title_no, job.title_en, job.position_title_no, job.position_title_en, job.description_short_no, job.description_short_en, job.description_long_no, job.description_long_en, job.job_type, job.time_publish, job.application_deadline, job.banner_image, job.organization, job.application_url, job.updated_at, job.created_at, job.deleted_at, org."shortname", org."name_no", org."name_en", 
		array(SELECT "skill" FROM "skill" WHERE "ad" = $1::int) AS skills,
		array(SELECT "city" FROM "ad_city_relation" WHERE "ad" = $1::int) AS cities
    FROM "job_advertisement" AS job
    INNER JOIN "organization" AS org ON job."organization" = org."shortname"
    WHERE job."id" = $1::int LIMIT 1
`

type GetJobRow struct {
	ID                  int32       `json:"id"`
	Visible             bool        `json:"visible"`
	TitleNo             string      `json:"title_no"`
	TitleEn             zero.String `json:"title_en"`
	PositionTitleNo     string      `json:"position_title_no"`
	PositionTitleEn     zero.String `json:"position_title_en"`
	DescriptionShortNo  string      `json:"description_short_no"`
	DescriptionShortEn  zero.String `json:"description_short_en"`
	DescriptionLongNo   string      `json:"description_long_no"`
	DescriptionLongEn   zero.String `json:"description_long_en"`
	JobType             JobType     `json:"job_type"`
	TimePublish         time.Time   `json:"time_publish"`
	ApplicationDeadline time.Time   `json:"application_deadline"`
	BannerImage         zero.String `json:"banner_image"`
	Organization        string      `json:"organization"`
	ApplicationUrl      string      `json:"application_url"`
	UpdatedAt           time.Time   `json:"updated_at"`
	CreatedAt           time.Time   `json:"created_at"`
	DeletedAt           zero.Time   `json:"deleted_at"`
	Shortname           string      `json:"shortname"`
	NameNo              string      `json:"name_no"`
	NameEn              zero.String `json:"name_en"`
	Skills              interface{} `json:"skills"`
	Cities              interface{} `json:"cities"`
}

func (q *Queries) GetJob(ctx context.Context, id int32) (GetJobRow, error) {
	row := q.db.QueryRowContext(ctx, getJob, id)
	var i GetJobRow
	err := row.Scan(
		&i.ID,
		&i.Visible,
		&i.TitleNo,
		&i.TitleEn,
		&i.PositionTitleNo,
		&i.PositionTitleEn,
		&i.DescriptionShortNo,
		&i.DescriptionShortEn,
		&i.DescriptionLongNo,
		&i.DescriptionLongEn,
		&i.JobType,
		&i.TimePublish,
		&i.ApplicationDeadline,
		&i.BannerImage,
		&i.Organization,
		&i.ApplicationUrl,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.DeletedAt,
		&i.Shortname,
		&i.NameNo,
		&i.NameEn,
		&i.Skills,
		&i.Cities,
	)
	return i, err
}

const getJobs = `-- name: GetJobs :many
SELECT job."id", job."title_no", job."title_en", 
        job."position_title_no", job."position_title_en",
        job."job_type", job."time_publish", job."application_deadline",
        job."application_url", job."updated_at", job."deleted_at", job."visible",
        org."name_no", org."name_en"
    FROM "job_advertisement" AS job
    INNER JOIN "organization" AS org ON job."organization" = org."id"
    ORDER BY job."id"
    LIMIT $2::int
    OFFSET $1::int
`

type GetJobsParams struct {
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
}

type GetJobsRow struct {
	ID                  int32       `json:"id"`
	TitleNo             string      `json:"title_no"`
	TitleEn             zero.String `json:"title_en"`
	PositionTitleNo     string      `json:"position_title_no"`
	PositionTitleEn     zero.String `json:"position_title_en"`
	JobType             JobType     `json:"job_type"`
	TimePublish         time.Time   `json:"time_publish"`
	ApplicationDeadline time.Time   `json:"application_deadline"`
	ApplicationUrl      string      `json:"application_url"`
	UpdatedAt           time.Time   `json:"updated_at"`
	DeletedAt           zero.Time   `json:"deleted_at"`
	Visible             bool        `json:"visible"`
	NameNo              string      `json:"name_no"`
	NameEn              zero.String `json:"name_en"`
}

func (q *Queries) GetJobs(ctx context.Context, arg GetJobsParams) ([]GetJobsRow, error) {
	rows, err := q.db.QueryContext(ctx, getJobs, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetJobsRow{}
	for rows.Next() {
		var i GetJobsRow
		if err := rows.Scan(
			&i.ID,
			&i.TitleNo,
			&i.TitleEn,
			&i.PositionTitleNo,
			&i.PositionTitleEn,
			&i.JobType,
			&i.TimePublish,
			&i.ApplicationDeadline,
			&i.ApplicationUrl,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.Visible,
			&i.NameNo,
			&i.NameEn,
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

const softDeleteJob = `-- name: SoftDeleteJob :one
UPDATE "job_advertisement"
SET
    "deleted_at" = now()
WHERE "id" = $1::int RETURNING id, visible, title_no, title_en, position_title_no, position_title_en, description_short_no, description_short_en, description_long_no, description_long_en, job_type, time_publish, application_deadline, banner_image, organization, application_url, updated_at, created_at, deleted_at
`

func (q *Queries) SoftDeleteJob(ctx context.Context, id int32) (JobAdvertisement, error) {
	row := q.db.QueryRowContext(ctx, softDeleteJob, id)
	var i JobAdvertisement
	err := row.Scan(
		&i.ID,
		&i.Visible,
		&i.TitleNo,
		&i.TitleEn,
		&i.PositionTitleNo,
		&i.PositionTitleEn,
		&i.DescriptionShortNo,
		&i.DescriptionShortEn,
		&i.DescriptionLongNo,
		&i.DescriptionLongEn,
		&i.JobType,
		&i.TimePublish,
		&i.ApplicationDeadline,
		&i.BannerImage,
		&i.Organization,
		&i.ApplicationUrl,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const updateJob = `-- name: UpdateJob :one
UPDATE "job_advertisement"
SET
    "visible" = COALESCE($1, visible),
    "title_no" = COALESCE($2, title_no),
    "title_en" = COALESCE($3, title_en),
    "position_title_no" = COALESCE($4, position_title_no),
    "position_title_en" = COALESCE($5, position_title_en),
    "description_short_no" = COALESCE($6, description_short_no),
    "description_short_en" = COALESCE($7, description_short_en),
    "description_long_no" = COALESCE($8, description_long_no),
    "description_long_en" = COALESCE($9, description_long_en),
    "job_type" = COALESCE($10, job_type),
    "time_publish" = COALESCE($11, time_publish),
    "application_deadline" = COALESCE($12, application_deadline),
    "banner_image" = COALESCE($13, banner_image),
    "organization" = COALESCE($14, organization),
    "application_url" = COALESCE($15, application_url),
    "updated_at" = now()
WHERE "id" = $16::int RETURNING id, visible, title_no, title_en, position_title_no, position_title_en, description_short_no, description_short_en, description_long_no, description_long_en, job_type, time_publish, application_deadline, banner_image, organization, application_url, updated_at, created_at, deleted_at
`

type UpdateJobParams struct {
	Visible             zero.Bool   `json:"visible"`
	TitleNo             zero.String `json:"title_no"`
	TitleEn             zero.String `json:"title_en"`
	PositionTitleNo     zero.String `json:"position_title_no"`
	PositionTitleEn     zero.String `json:"position_title_en"`
	DescriptionShortNo  zero.String `json:"description_short_no"`
	DescriptionShortEn  zero.String `json:"description_short_en"`
	DescriptionLongNo   zero.String `json:"description_long_no"`
	DescriptionLongEn   zero.String `json:"description_long_en"`
	JobType             NullJobType `json:"job_type"`
	TimePublish         zero.Time   `json:"time_publish"`
	ApplicationDeadline zero.Time   `json:"application_deadline"`
	BannerImage         zero.String `json:"banner_image"`
	Organization        zero.String `json:"organization"`
	ApplicationUrl      zero.String `json:"application_url"`
	ID                  int32       `json:"id"`
}

func (q *Queries) UpdateJob(ctx context.Context, arg UpdateJobParams) (JobAdvertisement, error) {
	row := q.db.QueryRowContext(ctx, updateJob,
		arg.Visible,
		arg.TitleNo,
		arg.TitleEn,
		arg.PositionTitleNo,
		arg.PositionTitleEn,
		arg.DescriptionShortNo,
		arg.DescriptionShortEn,
		arg.DescriptionLongNo,
		arg.DescriptionLongEn,
		arg.JobType,
		arg.TimePublish,
		arg.ApplicationDeadline,
		arg.BannerImage,
		arg.Organization,
		arg.ApplicationUrl,
		arg.ID,
	)
	var i JobAdvertisement
	err := row.Scan(
		&i.ID,
		&i.Visible,
		&i.TitleNo,
		&i.TitleEn,
		&i.PositionTitleNo,
		&i.PositionTitleEn,
		&i.DescriptionShortNo,
		&i.DescriptionShortEn,
		&i.DescriptionLongNo,
		&i.DescriptionLongEn,
		&i.JobType,
		&i.TimePublish,
		&i.ApplicationDeadline,
		&i.BannerImage,
		&i.Organization,
		&i.ApplicationUrl,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}