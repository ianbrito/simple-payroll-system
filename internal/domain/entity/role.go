package entity

import "github.com/google/uuid"

type Role struct {
	ID   string
	Name string
}

func NewRole(name string) *Role {
	return &Role{
		ID:   uuid.New().String(),
		Name: name,
	}
}
