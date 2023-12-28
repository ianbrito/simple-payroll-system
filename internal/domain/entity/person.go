package entity

import (
	"time"

	"github.com/google/uuid"
)

type Person struct {
	ID    string
	Name  string
	CPF   string
	Birth time.Time
}

func NewPerson(name string, cpf string, birth time.Time) *Person {
	return &Person{
		ID:    uuid.New().String(),
		Name:  name,
		CPF:   cpf,
		Birth: birth,
	}
}
