package service

import (
	"context"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	GenerateCookie(ctx context.Context, id int) (string, error)
	DeleteCookie(ctx context.Context, session string) error
	GetSessionValue(ctx context.Context, session string) (int, error)
	CreateUser(ctx context.Context, user model.User) (int, error)
	GetUser(ctx context.Context, mail, password string) (model.User, error)
}

type Feed interface {
	GetNextUser(ctx context.Context, userId int) (model.User, error)
}

type User interface {
	GetCurrentUser(ctx context.Context, userId int) (model.User, error)
	//TODO: edit user
}

type Service struct {
	Authorization
	Feed
	User
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.User, repo.Store),
		Feed:          NewFeedService(repo.User, repo.Store),
		User:          NewUserService(repo.User, repo.Store),
	}
}
