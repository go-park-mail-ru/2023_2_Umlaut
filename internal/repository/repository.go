package repository

import (
	"context"
	core2 "github.com/go-park-mail-ru/2023_2_Umlaut/internal/model/core"
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
	CreateUser(ctx context.Context, user core2.User) (int, error)
	GetUser(ctx context.Context, mail string) (core2.User, error)
	GetUserById(ctx context.Context, id int) (core2.User, error)
	GetNextUser(ctx context.Context, user core2.User, params core2.FilterParams) (core2.User, error)
	UpdateUser(ctx context.Context, user core2.User) (core2.User, error)
	UpdateUserPassword(ctx context.Context, user core2.User) error
	ShowCSAT(ctx context.Context, userId int) (bool, error)
	GetUserInvites(ctx context.Context, userId int) (int, error)
	ResetLikeCounter(ctx context.Context) error
}

type Like interface {
	CreateLike(ctx context.Context, like core2.Like) (core2.Like, error)
	IsMutualLike(ctx context.Context, like core2.Like) (bool, error)
	GetUserLikedToLikes(ctx context.Context, userId int) ([]core2.PremiumLike, error)
	ResetDislike(ctx context.Context) error
}

type Dialog interface {
	CreateDialog(ctx context.Context, dialog core2.Dialog) (int, error)
	GetDialogs(ctx context.Context, userId int) ([]core2.Dialog, error)
	GetDialogById(ctx context.Context, id int) (core2.Dialog, error)
}

type Message interface {
	GetDialogMessages(ctx context.Context, userId int, recipientId int) ([]core2.Message, error)
	CreateMessage(ctx context.Context, message core2.Message) (core2.Message, error)
	UpdateMessage(ctx context.Context, message core2.Message) (core2.Message, error)
}

type Tag interface {
	GetAllTags(ctx context.Context) ([]string, error)
}

type Admin interface {
	GetAdmin(ctx context.Context, mail string) (core2.Admin, error)
	CreateFeedback(ctx context.Context, stat core2.Feedback) (int, error)
	CreateRecommendation(ctx context.Context, rec core2.Recommendation) (int, error)
	CreateFeedFeedback(ctx context.Context, rec core2.Recommendation) (int, error)
	GetFeedbacks(ctx context.Context) ([]core2.Feedback, error)
	GetRecommendations(ctx context.Context) ([]core2.Recommendation, error)
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
	GetComplaintTypes(ctx context.Context) ([]core2.ComplaintType, error)
	CreateComplaint(ctx context.Context, complaint core2.Complaint) (int, error)
	GetNextComplaint(ctx context.Context) (core2.Complaint, error)
	DeleteComplaint(ctx context.Context, complaintId int) error
	AcceptComplaint(ctx context.Context, complaintId int) (core2.Complaint, error)
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
