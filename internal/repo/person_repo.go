package repo

import (
	"context"
	"fmt"

	"github.com/AlexeyLoychenko/person_api/internal/entity"
	"github.com/AlexeyLoychenko/person_api/internal/model"
	"github.com/AlexeyLoychenko/person_api/pkg/postgres"
	"github.com/jackc/pgx/v5"
)

type PersonRepo struct {
	c *postgres.Conn
}

func New(c *postgres.Conn) *PersonRepo {
	return &PersonRepo{c}
}

func (r *PersonRepo) Get(req model.GetPersonRequest) ([]entity.Person, error) {

	stmt := "select id, name, surname, patronymic, gender, age, nationality, created_at, updated_at from paginate($1, $2, $3, $4, $5, $6, $7, $8)"

	rows, err := r.c.Pool.Query(context.Background(), stmt, req.PageId, req.RecordsPerPage, req.Name, req.Surname, req.Patronymic, req.Gender, req.Age, req.Nationality)
	if err != nil {
		return []entity.Person{}, fmt.Errorf("person repo - Get - query(): %w", err)
	}
	defer rows.Close()

	res, err := pgx.CollectRows[entity.Person](rows, pgx.RowToStructByName[entity.Person])
	if err != nil {
		return []entity.Person{}, fmt.Errorf("person repo - Get - collectrows(): %w", err)
	}
	return res, nil
}

func (r *PersonRepo) GetById(id int) (entity.Person, error) {
	stmt := "select id, name, surname, patronymic, gender, age, nationality, created_at, updated_at from person where id=$1"
	rows, err := r.c.Pool.Query(context.Background(), stmt, id)
	if err != nil {
		return entity.Person{}, fmt.Errorf("person repo - GetById - Query(): %w", err)
	}
	defer rows.Close()

	var p entity.Person
	rowsCnt := 0
	for rows.Next() {
		err = rows.Scan(&p.Id, &p.Name, &p.Surname, &p.Patronymic, &p.Gender, &p.Age, &p.Nationality, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return entity.Person{}, fmt.Errorf("person repo - GetById - rows.Scan(): %w", err)
		}
		rowsCnt++
	}
	if rowsCnt == 0 {
		return entity.Person{}, ErrRecordNotExist
	}

	return p, nil
}

func (r *PersonRepo) Delete(id int) error {
	var recordId int
	stmt := "delete from person where id=$1 returning id"
	err := r.c.Pool.QueryRow(context.Background(), stmt, id).Scan(&recordId)
	if err != nil {
		if err == pgx.ErrNoRows {
			return ErrRecordNotExist
		}
		return fmt.Errorf("person repo - Delete - QueryRow(): %w", err)
	}

	return nil
}

func (r *PersonRepo) Update(req model.UpdatePersonRequest) error {
	var recordId int
	stmt := "update person " +
		"set name=$1, surname=$2, patronymic=$3, gender=$4, age=$5, nationality=$6, updated_at=current_timestamp " +
		"where id=$7 " +
		"returning id"
	err := r.c.Pool.QueryRow(context.Background(), stmt, req.Name, req.Surname, req.Patronymic, req.Gender, req.Age, req.Nationality, req.Id).Scan(&recordId)
	if err != nil {
		if err == pgx.ErrNoRows {
			return ErrRecordNotExist
		}
		return fmt.Errorf("person repo - Update - QueryRow(): %w", err)
	}
	return nil
}

func (r *PersonRepo) Create(req model.CreatePersonRequest) (entity.Person, error) {
	var recordId int
	stmt := "insert into person (name, surname, patronymic, gender, age, nationality) " +
		"values ($1,$2,$3,$4,$5,$6) " +
		"returning id"
	err := r.c.Pool.QueryRow(context.Background(), stmt, req.Name, req.Surname, req.Patronymic, req.Gender, req.Age, req.Nationality).Scan(&recordId)
	if err != nil {
		return entity.Person{}, fmt.Errorf("person repo - Create - QueryRow(): %w", err)
	}
	return r.GetById(recordId)
}
