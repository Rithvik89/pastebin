// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: queries.sql

package sqlc_db

import (
	"context"
	"time"
)

const createPaste = `-- name: CreatePaste :exec
insert into pastes (
  shortlink,
  expiration_length_in_minutes,
  created_at,
  paste_path
) Values (?,?,?,?)
`

type CreatePasteParams struct {
	Shortlink                 string
	ExpirationLengthInMinutes int32
	CreatedAt                 time.Time
	PastePath                 string
}

func (q *Queries) CreatePaste(ctx context.Context, arg CreatePasteParams) error {
	_, err := q.db.ExecContext(ctx, createPaste,
		arg.Shortlink,
		arg.ExpirationLengthInMinutes,
		arg.CreatedAt,
		arg.PastePath,
	)
	return err
}

const getPaste = `-- name: GetPaste :one
select paste_path,expiration_length_in_minutes,created_at from pastes
where shortlink = (?)
`

type GetPasteRow struct {
	PastePath                 string
	ExpirationLengthInMinutes int32
	CreatedAt                 time.Time
}

func (q *Queries) GetPaste(ctx context.Context, shortlink string) (GetPasteRow, error) {
	row := q.db.QueryRowContext(ctx, getPaste, shortlink)
	var i GetPasteRow
	err := row.Scan(&i.PastePath, &i.ExpirationLengthInMinutes, &i.CreatedAt)
	return i, err
}
