package service

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
	"net/http"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	GenerateCookie(ctx context.Context, id int) (*http.Cookie, error)
	DeleteCookie(ctx context.Context, session *http.Cookie) error
	CreateUser(user model.User) (int, error)
	GetUser(mail, password string) (model.User, error)
}

type Feed interface {
	GetNextUser(ctx context.Context, session *http.Cookie) (model.User, error)
}

type Service struct {
	Authorization
	Feed
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.User, repo.Store),
		Feed:          NewFeedService(repo.User, repo.Store),
	}
}
