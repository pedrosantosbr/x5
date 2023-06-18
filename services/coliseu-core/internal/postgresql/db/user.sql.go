// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: user.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const InsertUser = `-- name: InsertUser :one
INSERT INTO users (
  email,
  password
)
VALUES (
  $1,
  $2
)
RETURNING id, created_at, updated_at
`

type InsertUserParams struct {
	Email    string
	Password string
}

type InsertUserRow struct {
	ID        uuid.UUID
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) (InsertUserRow, error) {
	row := q.db.QueryRow(ctx, InsertUser, arg.Email, arg.Password)
	var i InsertUserRow
	err := row.Scan(&i.ID, &i.CreatedAt, &i.UpdatedAt)
	return i, err
}