package service

import (
	"context"
	"mime/multipart"

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
	GetNextUsers(ctx context.Context, userId int) ([]model.User, error)

}

type User interface {
	GetCurrentUser(ctx context.Context, userId int) (model.User, error)
	UpdateUser(ctx context.Context, user model.User) (model.User, error)
	UpdateUserPhoto(ctx context.Context, userId int, imagePath string) error
	CreateFile(ctx context.Context, userId int, file multipart.File, size int64) (string, error)
	GetFile(ctx context.Context, userId int, fileName string) ([]byte, string, error)
	DeleteFile(ctx context.Context, userId int, fileName string) error
}

type Like interface {
	CreateLike(ctx context.Context, like model.Like) error 
	IsUserLiked(ctx context.Context, like model.Like) (bool, error)
	IsLikeExists(ctx context.Context, like model.Like) (bool, error)
}

type Dialog interface {
	CreateDialog(ctx context.Context, dialog model.Dialog) (int, error)
	GetDialogs(ctx context.Context, userId int) ([]model.Dialog, error)
}

type Tag interface {
	GetAllTags(ctx context.Context) ([]model.Tag, error)
	GetUserTags(ctx context.Context, userId int) ([]model.Tag, error)
}

type Service struct {
	Authorization
	Feed
	User
	Like
	Dialog
	Tag
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.User, repo.Store),
		Feed:          NewFeedService(repo.User, repo.Store, repo.Dialog),
		User:          NewUserService(repo.User, repo.Store, repo.FileServer),
		Like:          NewLikeService(repo.Like),
		Dialog:        NewDialogService(repo.Dialog),
		Tag:           NewTagService(repo.Tag),
	}
}
