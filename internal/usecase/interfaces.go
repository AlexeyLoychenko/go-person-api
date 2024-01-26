package usecase

import (
	"github.com/AlexeyLoychenko/person_api/internal/entity"
	"github.com/AlexeyLoychenko/person_api/internal/model"
)

type PersonRepo interface {
	Get(model.GetPersonRequest) ([]entity.Person, error)
	GetById(int) (entity.Person, error)
	Delete(int) error
	Update(model.UpdatePersonRequest) error
	Create(model.CreatePersonRequest) (entity.Person, error)
}

type UseCase interface {
	GetById(id int) (entity.Person, error)
	GetList(req model.GetPersonRequest) ([]entity.Person, entity.PersonPage, error)
	Delete(id int) (bool, error)
	Update(req model.UpdatePersonRequest) (bool, error)
	Create(req model.CreatePersonRequest) (entity.Person, error)
}

type PersonApi interface {
	CollectPersonData(string) (model.GroupedApiResponse, error)
}
