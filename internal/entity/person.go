package entity

import (
	"database/sql"
)

type Person struct {
	Id          int          `json:"id" db:"id"`
	Name        string       `json:"name" db:"name"`
	Surname     string       `json:"surname" db:"surname"`
	Patronymic  string       `json:"patronymic" db:"patronymic"`
	Gender      string       `json:"gender" db:"gender"`
	Age         int          `json:"age" db:"age"`
	Nationality string       `json:"nationality" db:"nationality"`
	CreatedAt   sql.NullTime `json:"created_at" db:"created_at"`
	UpdatedAt   sql.NullTime `json:"updated_at" db:"updated_at"`
}

type PersonPage struct {
	PageId      int  `json:"page_id"`
	HasNextPage bool `json:"has_next_page"`
}
