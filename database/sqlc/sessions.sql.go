// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: sessions.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createSession = `-- name: CreateSession :one
INSERT INTO sessions (
  user_id, "access_token", "refresh_token", "expires_at"
) VALUES (
  $1, $2, $3, $4
)
RETURNING id, user_id, access_token, refresh_token, expires_at, created_at
`

type CreateSessionParams struct {
	UserID       pgtype.UUID
	AccessToken  string
	RefreshToken string
	ExpiresAt    int64
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error) {
	row := q.db.QueryRow(ctx, createSession,
		arg.UserID,
		arg.AccessToken,
		arg.RefreshToken,
		arg.ExpiresAt,
	)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.AccessToken,
		&i.RefreshToken,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}

const deleteSession = `-- name: DeleteSession :exec
DELETE FROM sessions
WHERE user_id = $1
RETURNING id, user_id, access_token, refresh_token, expires_at, created_at
`

func (q *Queries) DeleteSession(ctx context.Context, userID pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteSession, userID)
	return err
}

const getSession = `-- name: GetSession :one
SELECT id, user_id, access_token, refresh_token, expires_at, created_at FROM sessions
WHERE user_id = $1 LIMIT 1
`

func (q *Queries) GetSession(ctx context.Context, userID pgtype.UUID) (Session, error) {
	row := q.db.QueryRow(ctx, getSession, userID)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.AccessToken,
		&i.RefreshToken,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}

const updateSession = `-- name: UpdateSession :one
UPDATE sessions
SET "access_token" = $2, "refresh_token" = $3, "expires_at" = $4
WHERE user_id = $1
RETURNING id, user_id, access_token, refresh_token, expires_at, created_at
`

type UpdateSessionParams struct {
	UserID       pgtype.UUID
	AccessToken  string
	RefreshToken string
	ExpiresAt    int64
}

func (q *Queries) UpdateSession(ctx context.Context, arg UpdateSessionParams) (Session, error) {
	row := q.db.QueryRow(ctx, updateSession,
		arg.UserID,
		arg.AccessToken,
		arg.RefreshToken,
		arg.ExpiresAt,
	)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.AccessToken,
		&i.RefreshToken,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}
