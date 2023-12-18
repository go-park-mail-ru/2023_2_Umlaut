package service

import (
	"context"
	core2 "github.com/go-park-mail-ru/2023_2_Umlaut/internal/model/core"
	"mime/multipart"

	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	GenerateCookie(ctx context.Context, id int) (string, error)
	DeleteCookie(ctx context.Context, session string) error
	GetSessionValue(ctx context.Context, session string) (int, error)
	CreateUser(ctx context.Context, user core2.User) (int, error)
	GetUser(ctx context.Context, mail, password string) (core2.User, error)
	GetDecodeUserId(ctx context.Context, message string) (int, error)
}

type Feed interface {
	GetNextUser(ctx context.Context, params core2.FilterParams) (core2.FeedData, error)
}

type User interface {
	GetCurrentUser(ctx context.Context, userId int) (core2.User, error)
	UpdateUser(ctx context.Context, user core2.User) (core2.User, error)
	CreateFile(ctx context.Context, userId int, file multipart.File, size int64) (string, error)
	DeleteFile(ctx context.Context, userId int, link string) error
	GetUserShareCridentials(ctx context.Context, userId int) (int, string, error)
}

type Like interface {
	CreateLike(ctx context.Context, like core2.Like) (core2.Dialog, error)
	GetUserLikedToLikes(ctx context.Context, userId int) (bool, []core2.PremiumLike, error)
}

type Dialog interface {
	CreateDialog(ctx context.Context, dialog core2.Dialog) (int, error)
	GetDialogs(ctx context.Context, userId int) ([]core2.Dialog, error)
	GetDialog(ctx context.Context, id int) (core2.Dialog, error)
}

type Message interface {
	GetDialogMessages(ctx context.Context, userId int, recipientId int) ([]core2.Message, error)
	SaveOrUpdateMessage(ctx context.Context, message core2.Message) (core2.Message, error)
}

type Tag interface {
	GetAllTags(ctx context.Context) ([]string, error)
}

type Admin interface {
	CreateFeedback(ctx context.Context, stat core2.Feedback) (int, error)
	CreateRecommendation(ctx context.Context, rec core2.Recommendation) (int, error)
	CreateFeedFeedback(ctx context.Context, rec core2.Recommendation) (int, error)
	GetRecommendationsStatistics(ctx context.Context) (core2.RecommendationStatistic, error)
	GetFeedbackStatistics(ctx context.Context) (core2.FeedbackStatistic, error)
	GetCSATType(ctx context.Context, userId int) (int, error)
}

type Complaint interface {
	GetComplaintTypes(ctx context.Context) ([]core2.ComplaintType, error)
	CreateComplaint(ctx context.Context, complaint core2.Complaint) (int, error)
	GetNextComplaint(ctx context.Context) (core2.Complaint, error)
}

type Background interface {
	ResetLikeCounter(ctx context.Context) error
	ResetDislike(ctx context.Context) error
}

type Service struct {
	Authorization Authorization
	Feed          Feed
	User          User
	Like          Like
	Dialog        Dialog
	Message       Message
	Tag           Tag
	Admin         Admin
	Complaint     Complaint
	Background    Background
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.User, repo.Store, repo.Admin),
		Feed:          NewFeedService(repo.User, repo.Store, repo.Dialog),
		User:          NewUserService(repo.User, repo.Store, repo.FileServer),
		Like:          NewLikeService(repo.Like, repo.Dialog, repo.User),
		Dialog:        NewDialogService(repo.Dialog),
		Message:       NewMessageService(repo.Message),
		Tag:           NewTagService(repo.Tag),
		Admin:         NewAdminService(repo.Admin, repo.User),
		Complaint:     NewComplaintService(repo.Complaint),
		Background:    NewBackgroundService(repo.User, repo.Like),
	}
}
