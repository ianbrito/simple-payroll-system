package entity

import (
	"time"

	"github.com/google/uuid"
)

type Lotation struct {
	ID            string
	AdmissionAt   time.Time
	ResignationAt time.Time
	Employee      *Employee
	Department    *Department
	Role          *Role
	Bond          *Bond
}

func NewLotation(admissionAt time.Time, employee *Employee, department *Department, role *Role, bond *Bond) *Lotation {
	return &Lotation{
		ID:            uuid.New().String(),
		AdmissionAt:   admissionAt,
		ResignationAt: time.Time{},
		Employee:      employee,
		Department:    department,
		Role:          role,
		Bond:          bond,
	}
}
