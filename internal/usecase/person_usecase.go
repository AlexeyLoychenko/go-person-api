package usecase

import (
	"errors"
	"fmt"

	"github.com/AlexeyLoychenko/person_api/internal/entity"
	"github.com/AlexeyLoychenko/person_api/internal/model"
	"github.com/AlexeyLoychenko/person_api/internal/repo"
)

const entitiesPerPage = 5

type PersonUseCase struct {
	repo   PersonRepo
	webapi PersonApi
}

func New(r PersonRepo, w PersonApi) *PersonUseCase {
	return &PersonUseCase{repo: r, webapi: w}
}

func (uc *PersonUseCase) GetById(id int) (entity.Person, error) {
	data, err := uc.repo.GetById(id)

	if err != nil {
		if errors.Is(err, repo.ErrRecordNotExist) {
			return entity.Person{}, err
		}
		return entity.Person{}, fmt.Errorf("personUseCase - uc.repo.GetById(): %w", err)
	}

	return data, nil
}

func (uc *PersonUseCase) GetList(r model.GetPersonRequest) ([]entity.Person, entity.PersonPage, error) {
	if r.PageId == 0 {
		r.PageId = 1
	}
	r.RecordsPerPage = entitiesPerPage
	data, err := uc.repo.Get(r)
	if err != nil {
		return nil, entity.PersonPage{}, fmt.Errorf("personUseCase - uc.repo.GetPeople(): %w", err)
	}
	page := entity.PersonPage{PageId: r.PageId}
	if len(data) > entitiesPerPage {
		page.HasNextPage = true
		data = data[:entitiesPerPage]
	}

	return data, page, nil
}

func (uc *PersonUseCase) Delete(id int) (bool, error) {
	err := uc.repo.Delete(id)
	if err != nil {
		if errors.Is(err, repo.ErrRecordNotExist) {
			return false, err
		}
		return false, fmt.Errorf("personUseCase - uc.repo.Delete(): %w", err)
	}
	return true, nil
}

func (uc *PersonUseCase) Update(r model.UpdatePersonRequest) (bool, error) {
	err := uc.repo.Update(r)
	if err != nil {
		if errors.Is(err, repo.ErrRecordNotExist) {
			return false, err
		}
		return false, fmt.Errorf("personUseCase - uc.repo.Update(): %w", err)
	}
	return true, nil
}

func (uc *PersonUseCase) Create(r model.CreatePersonRequest) (entity.Person, error) {
	data, err := uc.webapi.CollectPersonData(r.Name)
	if err != nil {
		return entity.Person{}, fmt.Errorf("personUseCase - uc.webapi.CollectPersonData: %w", err)
	}
	r.Age = data.Age
	r.Gender = data.Gender
	r.Nationality = data.Nationality
	res, err := uc.repo.Create(r)
	if err != nil {
		return entity.Person{}, fmt.Errorf("personUseCase - uc.repo.Create(): %w", err)
	}
	return res, nil
}
