package model

import (
	"github.com/sh-tatsuno/shabm/db"
	"github.com/sh-tatsuno/shabm/entity"
)

type PersonModelInterface interface {
	GetAllPeople() ([]entity.Person, error)
	CreatePerson(person entity.Person) (uint, error)
}

type PersonModel struct {
}

func (m *PersonModel) GetAllPeople() ([]entity.Person, error) {
	var people []entity.Person
	dbc := db.Instance.Tx.Find(&people)
	return people, dbc.Error
}

func (m *PersonModel) CreatePerson(person entity.Person) (uint, error) {
	dbc := db.Instance.Tx.Create(&person)
	return person.ID, dbc.Error
}
