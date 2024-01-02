// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: lotations.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const deleteLotationByID = `-- name: DeleteLotationByID :exec
DELETE FROM lotations WHERE id = $1
`

func (q *Queries) DeleteLotationByID(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteLotationByID, id)
	return err
}

const getLotationByID = `-- name: GetLotationByID :one
SELECT id, admission_at, resignation_at, employee_id, department_id, role_id, bond_id, created_at, updated_at FROM lotations WHERE id = $1
`

func (q *Queries) GetLotationByID(ctx context.Context, id string) (Lotation, error) {
	row := q.db.QueryRowContext(ctx, getLotationByID, id)
	var i Lotation
	err := row.Scan(
		&i.ID,
		&i.AdmissionAt,
		&i.ResignationAt,
		&i.EmployeeID,
		&i.DepartmentID,
		&i.RoleID,
		&i.BondID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getLotations = `-- name: GetLotations :many
SELECT id, admission_at, resignation_at, employee_id, department_id, role_id, bond_id, created_at, updated_at FROM lotations ORDER BY created_at
`

func (q *Queries) GetLotations(ctx context.Context) ([]Lotation, error) {
	rows, err := q.db.QueryContext(ctx, getLotations)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Lotation
	for rows.Next() {
		var i Lotation
		if err := rows.Scan(
			&i.ID,
			&i.AdmissionAt,
			&i.ResignationAt,
			&i.EmployeeID,
			&i.DepartmentID,
			&i.RoleID,
			&i.BondID,
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

const insertLotation = `-- name: InsertLotation :exec
INSERT INTO lotations (id, admission_at, resignation_at, employee_id, department_id, role_id, bond_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
`

type InsertLotationParams struct {
	ID            string
	AdmissionAt   time.Time
	ResignationAt sql.NullTime
	EmployeeID    string
	DepartmentID  string
	RoleID        string
	BondID        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (q *Queries) InsertLotation(ctx context.Context, arg InsertLotationParams) error {
	_, err := q.db.ExecContext(ctx, insertLotation,
		arg.ID,
		arg.AdmissionAt,
		arg.ResignationAt,
		arg.EmployeeID,
		arg.DepartmentID,
		arg.RoleID,
		arg.BondID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}