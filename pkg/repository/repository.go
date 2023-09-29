package repository

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/redis/go-redis/v9"
)

type User interface {
	CreateUser(user model.User) (int, error)
	GetUser(mail string) (model.User, error)
	GetUserById(id int) (model.User, error)
	GetNextUser(user model.User) (model.User, error)
}

type Store interface {
	SetSession(ctx context.Context, SID string, id int, lifetime time.Duration) error
	GetSession(ctx context.Context, SID string) (int, error)
	DeleteSession(ctx context.Context, SID string) error
}

type Repository struct {
	User
	Store
}

func NewRepository(db *sqlx.DB, client *redis.Client) *Repository {
	return &Repository{
		User:  NewUserPostgres(db),
		Store: NewRedisStore(client),
	}
}
