package entity

import "github.com/google/uuid"

type Department struct {
	ID      string
	Name    string
	Acronym string
}

func NewDepartment(name string, acronym string) *Department {
	return &Department{
		ID:      uuid.New().String(),
		Name:    name,
		Acronym: acronym,
	}
}
