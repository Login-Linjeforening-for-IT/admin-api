// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: rules.sql

package db

import (
	"context"
	"time"

	"github.com/guregu/null/zero"
)

const createRule = `-- name: CreateRule :one
INSERT INTO "rule" (
    "name_no", "name_en",
    "description_no", "description_en"
) VALUES ($1, $2, $3, $4) RETURNING id, name_no, name_en, description_no, description_en, updated_at, created_at, deleted_at
`

type CreateRuleParams struct {
	NameNo        string      `json:"name_no"`
	NameEn        zero.String `json:"name_en"`
	DescriptionNo string      `json:"description_no"`
	DescriptionEn zero.String `json:"description_en"`
}

func (q *Queries) CreateRule(ctx context.Context, arg CreateRuleParams) (Rule, error) {
	row := q.db.QueryRowContext(ctx, createRule,
		arg.NameNo,
		arg.NameEn,
		arg.DescriptionNo,
		arg.DescriptionEn,
	)
	var i Rule
	err := row.Scan(
		&i.ID,
		&i.NameNo,
		&i.NameEn,
		&i.DescriptionNo,
		&i.DescriptionEn,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getRule = `-- name: GetRule :one
SELECT id, name_no, name_en, description_no, description_en, updated_at, created_at, deleted_at FROM "rule" WHERE "id" = $1::int LIMIT 1
`

func (q *Queries) GetRule(ctx context.Context, id int32) (Rule, error) {
	row := q.db.QueryRowContext(ctx, getRule, id)
	var i Rule
	err := row.Scan(
		&i.ID,
		&i.NameNo,
		&i.NameEn,
		&i.DescriptionNo,
		&i.DescriptionEn,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getRules = `-- name: GetRules :many
SELECT "id", "name_no", "name_en", "updated_at" FROM "rule"
    LIMIT $2::int
    OFFSET $1::int
`

type GetRulesParams struct {
	Offset int32 `json:"offset"`
	Limit  int32 `json:"limit"`
}

type GetRulesRow struct {
	ID        int32       `json:"id"`
	NameNo    string      `json:"name_no"`
	NameEn    zero.String `json:"name_en"`
	UpdatedAt time.Time   `json:"updated_at"`
}

func (q *Queries) GetRules(ctx context.Context, arg GetRulesParams) ([]GetRulesRow, error) {
	rows, err := q.db.QueryContext(ctx, getRules, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetRulesRow{}
	for rows.Next() {
		var i GetRulesRow
		if err := rows.Scan(
			&i.ID,
			&i.NameNo,
			&i.NameEn,
			&i.UpdatedAt,
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

const updateRule = `-- name: UpdateRule :one
UPDATE "rule"
SET
    "name_no" = COALESCE($1, name_no),
    "name_en" = COALESCE($2, name_en),
    "description_no" = COALESCE($3, description_no),
    "description_en" = COALESCE($4, description_en),
    "updated_at" = now()
WHERE "id" = $5::int
RETURNING id, name_no, name_en, description_no, description_en, updated_at, created_at, deleted_at
`

type UpdateRuleParams struct {
	NameNo        zero.String `json:"name_no"`
	NameEn        zero.String `json:"name_en"`
	DescriptionNo zero.String `json:"description_no"`
	DescriptionEn zero.String `json:"description_en"`
	ID            int32       `json:"id"`
}

func (q *Queries) UpdateRule(ctx context.Context, arg UpdateRuleParams) (Rule, error) {
	row := q.db.QueryRowContext(ctx, updateRule,
		arg.NameNo,
		arg.NameEn,
		arg.DescriptionNo,
		arg.DescriptionEn,
		arg.ID,
	)
	var i Rule
	err := row.Scan(
		&i.ID,
		&i.NameNo,
		&i.NameEn,
		&i.DescriptionNo,
		&i.DescriptionEn,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.DeletedAt,
	)
	return i, err
}
