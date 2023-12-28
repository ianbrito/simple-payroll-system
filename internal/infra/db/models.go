// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package db

import (
	"database/sql"
	"time"
)

type Bond struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Department struct {
	ID        string
	Name      string
	Acronym   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Employee struct {
	ID         string
	Enrollment string
	PersonID   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Lotation struct {
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

type Person struct {
	ID        string
	Name      string
	Cpf       string
	Birth     time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Role struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}