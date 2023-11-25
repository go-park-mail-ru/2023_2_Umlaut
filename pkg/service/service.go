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
	CreateFile(ctx context.Context, userId int, file multipart.File, size int64) (string, error)
	DeleteFile(ctx context.Context, userId int, link string) error
}

type Like interface {
	CreateLike(ctx context.Context, like model.Like) error
}

type Dialog interface {
	CreateDialog(ctx context.Context, dialog model.Dialog) (int, error)
	GetDialogs(ctx context.Context, userId int) ([]model.Dialog, error)
	GetDialogMessages(ctx context.Context, dialogId int) ([]model.Message, error)
}

type Tag interface {
	GetAllTags(ctx context.Context) ([]string, error)
}

type Complaint interface {
	GetComplaintTypes(ctx context.Context) ([]model.ComplaintType, error)
}

type Service struct {
	Authorization Authorization
	Feed          Feed
	User          User
	Like          Like
	Dialog        Dialog
	Tag           Tag
	Complaint     Complaint
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.User, repo.Store),
		Feed:          NewFeedService(repo.User, repo.Store, repo.Dialog),
		User:          NewUserService(repo.User, repo.Store, repo.FileServer),
		Like:          NewLikeService(repo.Like, repo.Dialog),
		Dialog:        NewDialogService(repo.Dialog),
		Tag:           NewTagService(repo.Tag),
		Complaint:     NewComplaintService(repo.Complaint),
	}
}
