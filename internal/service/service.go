package service

import (
	"context"
	"mime/multipart"

	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/model/core"
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/model/dto"

	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	GenerateCookie(ctx context.Context, id int) (string, error)
	DeleteCookie(ctx context.Context, session string) error
	GetSessionValue(ctx context.Context, session string) (int, error)
	CreateUser(ctx context.Context, user core.User) (int, error)
	GetUser(ctx context.Context, mail, password string) (core.User, error)
	GetDecodeUserId(ctx context.Context, message string) (int, error)
}

type Feed interface {
	GetNextUser(ctx context.Context, params dto.FilterParams) (dto.FeedData, error)
}

type User interface {
	GetCurrentUser(ctx context.Context, userId int) (core.User, error)
	UpdateUser(ctx context.Context, user core.User) (core.User, error)
	CreateFile(ctx context.Context, userId int, file multipart.File, size int64) (string, error)
	DeleteFile(ctx context.Context, userId int, link string) error
	GetUserShareCridentials(ctx context.Context, userId int) (int, string, error)
}

type Like interface {
	CreateLike(ctx context.Context, like core.Like) (core.Dialog, error)
	GetUserLikedToLikes(ctx context.Context, userId int) (bool, []dto.PremiumLike, error)
}

type Dialog interface {
	CreateDialog(ctx context.Context, dialog core.Dialog) (int, error)
	GetDialogs(ctx context.Context, userId int) ([]core.Dialog, error)
	GetDialog(ctx context.Context, id int) (core.Dialog, error)
}

type Message interface {
	GetDialogMessages(ctx context.Context, userId int, recipientId int) ([]core.Message, error)
	SaveOrUpdateMessage(ctx context.Context, message core.Message) (core.Message, error)
}

type Tag interface {
	GetAllTags(ctx context.Context) ([]string, error)
}

type Admin interface {
	CreateFeedback(ctx context.Context, stat core.Feedback) (int, error)
	CreateRecommendation(ctx context.Context, rec core.Recommendation) (int, error)
	CreateFeedFeedback(ctx context.Context, rec core.Recommendation) (int, error)
	GetRecommendationsStatistics(ctx context.Context) (core.RecommendationStatistic, error)
	GetFeedbackStatistics(ctx context.Context) (core.FeedbackStatistic, error)
	GetCSATType(ctx context.Context, userId int) (int, error)
}

type Complaint interface {
	GetComplaintTypes(ctx context.Context) ([]core.ComplaintType, error)
	CreateComplaint(ctx context.Context, complaint core.Complaint) (int, error)
	GetNextComplaint(ctx context.Context) (core.Complaint, error)
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
