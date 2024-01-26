package model

import "github.com/AlexeyLoychenko/person_api/internal/entity"

func NewPersonResponse(p entity.Person) *PersonResponse {
	return &PersonResponse{
		Id:          p.Id,
		Name:        p.Name,
		Surname:     p.Surname,
		Patronymic:  p.Patronymic,
		Gender:      p.Gender,
		Age:         p.Age,
		Nationality: p.Nationality,
		CreatedAt:   p.CreatedAt.Time.String(),
		UpdatedAt:   p.UpdatedAt.Time.String(),
	}
}

type PersonResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Gender      string `json:"gender"`
	Age         int    `json:"age"`
	Nationality string `json:"nationality"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type GetPersonRequest struct {
	Id             string
	Name           string
	Surname        string
	Patronymic     string
	Gender         string
	Age            int
	Nationality    string
	PageId         int
	RecordsPerPage int
}

type UpdatePersonRequest struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Gender      string `json:"gender"`
	Age         int    `json:"age"`
	Nationality string `json:"nationality"`
}

type CreatePersonRequest struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Gender      string `json:"gender"`
	Age         int    `json:"age"`
	Nationality string `json:"nationality"`
}
