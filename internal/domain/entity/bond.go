package entity

import "github.com/google/uuid"

type Bond struct {
	ID   string
	Name string
}

func NewBond(name string) *Bond {
	return &Bond{
		ID:   uuid.New().String(),
		Name: name,
	}
}
