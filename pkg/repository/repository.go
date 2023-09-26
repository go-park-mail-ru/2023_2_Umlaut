package repository

import (
	"database/sql"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/redis/go-redis/v9"
)

type User interface {
	CreateUser(user model.User) (int, error)
	GetUser(username, password string) (model.User, error)
}

type Repository struct {
	User
	SessionStore *redis.Client
}

func NewRepository(db *sql.DB, store *redis.Client) *Repository {
	return &Repository{
		User:         NewUserPostgres(db),
		SessionStore: store,
	}
}
