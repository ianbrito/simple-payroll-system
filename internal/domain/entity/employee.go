package entity

import "github.com/google/uuid"

type Employee struct {
	ID         string
	Enrollment string
	Person     *Person
}

func NewEmployee(enrollment string, person *Person) *Employee {
	return &Employee{
		ID:         uuid.New().String(),
		Enrollment: enrollment,
		Person:     person,
	}
}
