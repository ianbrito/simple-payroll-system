// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: roles.sql

package db

import (
	"context"
	"time"
)

const deleteRoleByID = `-- name: DeleteRoleByID :exec
DELETE FROM roles WHERE id = $1
`

func (q *Queries) DeleteRoleByID(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteRoleByID, id)
	return err
}

const getRoleByFields = `-- name: GetRoleByFields :one
SELECT id, name, created_at, updated_at FROM roles WHERE name = $1
`

func (q *Queries) GetRoleByFields(ctx context.Context, name string) (Role, error) {
	row := q.db.QueryRowContext(ctx, getRoleByFields, name)
	var i Role
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getRoleByID = `-- name: GetRoleByID :one
SELECT id, name, created_at, updated_at FROM roles WHERE id = $1
`

func (q *Queries) GetRoleByID(ctx context.Context, id string) (Role, error) {
	row := q.db.QueryRowContext(ctx, getRoleByID, id)
	var i Role
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getRoles = `-- name: GetRoles :many
SELECT id, name, created_at, updated_at FROM roles ORDER BY created_at
`

func (q *Queries) GetRoles(ctx context.Context) ([]Role, error) {
	rows, err := q.db.QueryContext(ctx, getRoles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Role
	for rows.Next() {
		var i Role
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
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

const insertRole = `-- name: InsertRole :one
INSERT INTO roles (id, name, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id, name, created_at, updated_at
`

type InsertRoleParams struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) InsertRole(ctx context.Context, arg InsertRoleParams) (Role, error) {
	row := q.db.QueryRowContext(ctx, insertRole,
		arg.ID,
		arg.Name,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Role
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
