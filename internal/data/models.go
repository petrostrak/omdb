package data

import (
	"database/sql"
	"errors"
)

var ErrRecordNotFound = errors.New("record not Found")

type Models struct {
	Movies MovieModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
	}
}