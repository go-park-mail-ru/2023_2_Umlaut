package repository

import (
	"database/sql"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
)

type UserPostgres struct {
	db *sql.DB
}

func NewUserPostgres(db *sql.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) CreateUser(user model.User) (int, error) {
	return 0, nil
}

func (r *UserPostgres) GetUser(username, password string) (model.User, error) {
	return model.User{}, nil
}
