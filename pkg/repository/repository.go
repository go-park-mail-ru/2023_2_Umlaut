package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/redis/go-redis/v9"
)

type User interface {
	CreateUser(ctx context.Context, user model.User) (int, error)
	GetUser(ctx context.Context, mail string) (model.User, error)
	GetUserById(ctx context.Context, id int) (model.User, error)
	GetNextUser(ctx context.Context, user model.User) (model.User, error)
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

func NewRepository(db *sql.DB, client *redis.Client) *Repository {
	return &Repository{
		User:  NewUserPostgres(db),
		Store: NewRedisStore(client),
	}
}
