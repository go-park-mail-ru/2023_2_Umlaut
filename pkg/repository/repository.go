package repository

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/core"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/dto"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"io"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type User interface {
	CreateUser(ctx context.Context, user core.User) (int, error)
	InsertOrUpdateUser(ctx context.Context, user core.User) (int, error)
	GetUser(ctx context.Context, mail string) (core.User, error)
	GetUserById(ctx context.Context, id int) (core.User, error)
	GetNextUser(ctx context.Context, user core.User, params dto.FilterParams) (core.User, error)
	UpdateUser(ctx context.Context, user core.User) (core.User, error)
	UpdateUserPhoto(ctx context.Context, user core.User) error
	UpdateUserPassword(ctx context.Context, user core.User) error
	ShowCSAT(ctx context.Context, userId int) (bool, error)
	GetUserInvites(ctx context.Context, userId int) (int, error)
	ResetLikeCounter(ctx context.Context) error
}

type Like interface {
	CreateLike(ctx context.Context, like core.Like) (core.Like, error)
	IsMutualLike(ctx context.Context, like core.Like) (bool, error)
	GetUserLikedToLikes(ctx context.Context, userId int) ([]dto.PremiumLike, error)
	ResetDislike(ctx context.Context) error
}

type Dialog interface {
	CreateDialog(ctx context.Context, dialog core.Dialog) (int, error)
	GetDialogs(ctx context.Context, userId int) ([]core.Dialog, error)
	GetDialogById(ctx context.Context, id int, userId int) (core.Dialog, error)
}

type Message interface {
	GetDialogMessages(ctx context.Context, userId int, recipientId int) ([]core.Message, error)
	CreateMessage(ctx context.Context, message core.Message) (core.Message, error)
	UpdateMessage(ctx context.Context, message core.Message) (core.Message, error)
}

type Tag interface {
	GetAllTags(ctx context.Context) ([]string, error)
}

type Admin interface {
	GetAdmin(ctx context.Context, mail string) (core.Admin, error)
	CreateFeedback(ctx context.Context, stat core.Feedback) (int, error)
	CreateRecommendation(ctx context.Context, rec core.Recommendation) (int, error)
	CreateFeedFeedback(ctx context.Context, rec core.Recommendation) (int, error)
	GetFeedbacks(ctx context.Context) ([]core.Feedback, error)
	GetRecommendations(ctx context.Context) ([]core.Recommendation, error)
	ShowFeedback(ctx context.Context, userId int) (bool, error)
	ShowRecommendation(ctx context.Context, userId int) (bool, error)
}

type Store interface {
	SetSession(ctx context.Context, SID string, id int, lifetime time.Duration) error
	GetSession(ctx context.Context, SID string) (int, error)
	DeleteSession(ctx context.Context, SID string) error
}

type FileServer interface {
	UploadFile(ctx context.Context, bucketName, fileName, contentType string, file io.Reader, size int64) (string, error)
	DeleteFile(ctx context.Context, bucketName, link string) error
	CreateBucket(ctx context.Context, bucketName string) error
}

type Complaint interface {
	GetComplaintTypes(ctx context.Context) ([]core.ComplaintType, error)
	CreateComplaint(ctx context.Context, complaint core.Complaint) (int, error)
	GetNextComplaint(ctx context.Context) (core.Complaint, error)
	DeleteComplaint(ctx context.Context, complaintId int) error
	AcceptComplaint(ctx context.Context, complaintId int) (core.Complaint, error)
}

type Repository struct {
	User       User
	Like       Like
	Dialog     Dialog
	Message    Message
	Tag        Tag
	Admin      Admin
	Store      Store
	FileServer FileServer
	Complaint  Complaint
}

type PgxPoolInterface interface {
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)

	Begin(context.Context) (pgx.Tx, error)
	Ping(ctx context.Context) error
	Close()
}

func NewRepository(db *pgxpool.Pool, dbAdmin *pgxpool.Pool, redisClient *redis.Client, minioClient *minio.Client) *Repository {
	return &Repository{
		User:       NewUserPostgres(db),
		Like:       NewLikePostgres(db),
		Dialog:     NewDialogPostgres(db),
		Message:    NewMessagePostgres(db),
		Tag:        NewTagPostgres(db),
		Admin:      NewAdminPostgres(dbAdmin),
		Complaint:  NewComplaintPostgres(db),
		Store:      NewRedisStore(redisClient),
		FileServer: NewMinioProvider(minioClient),
	}
}
