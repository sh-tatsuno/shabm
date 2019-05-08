package service

import (
	"github.com/sh-tatsuno/shabm/entity"
	"github.com/sh-tatsuno/shabm/model"
)

type PersonServiceInterface interface {
	GetAllPeople() ([]entity.Person, error)
	CreatePerson(person entity.Person) (uint, error)
}

type PersonService struct {
	PersonModel model.PersonModelInterface
}

func NewPersonService() PersonService {
	return PersonService{
		PersonModel: &model.PersonModel{},
	}
}

func (s *PersonService) GetAllPeople() ([]entity.Person, error) {
	bookmarks, err := s.PersonModel.GetAllPeople()
	return bookmarks, err
}

func (s *PersonService) CreatePerson(person entity.Person) (uint, error) {
	id, err := s.PersonModel.CreatePerson(person)
	return id, err
}
