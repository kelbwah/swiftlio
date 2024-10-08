// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: one_time_codes.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createOneTimeCode = `-- name: CreateOneTimeCode :one
INSERT INTO one_time_codes (id, user_id)
VALUES ($1, $2)
RETURNING id, code, expires_at, user_id
`

type CreateOneTimeCodeParams struct {
	ID     uuid.UUID
	UserID uuid.UUID
}

func (q *Queries) CreateOneTimeCode(ctx context.Context, arg CreateOneTimeCodeParams) (OneTimeCode, error) {
	row := q.db.QueryRowContext(ctx, createOneTimeCode, arg.ID, arg.UserID)
	var i OneTimeCode
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.ExpiresAt,
		&i.UserID,
	)
	return i, err
}

const getOneTimeCode = `-- name: GetOneTimeCode :one
SELECT code FROM one_time_codes
WHERE code = $1 and user_id = $2
`

type GetOneTimeCodeParams struct {
	Code   int32
	UserID uuid.UUID
}

func (q *Queries) GetOneTimeCode(ctx context.Context, arg GetOneTimeCodeParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, getOneTimeCode, arg.Code, arg.UserID)
	var code int32
	err := row.Scan(&code)
	return code, err
}
