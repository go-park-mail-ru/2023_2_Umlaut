package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"io"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
)

//go:generate mockgen -source=repository.go -destination=mocks/mock.go

type User interface {
	CreateUser(ctx context.Context, user model.User) (int, error)
	GetUser(ctx context.Context, mail string) (model.User, error)
	GetUserById(ctx context.Context, id int) (model.User, error)
	GetNextUser(ctx context.Context, user model.User, params model.FilterParams) (model.User, error)
	UpdateUser(ctx context.Context, user model.User) (model.User, error)
	UpdateUserPassword(ctx context.Context, user model.User) error
	ShowCSAT(ctx context.Context, userId int) (bool, error)
}

type Like interface {
	CreateLike(ctx context.Context, like model.Like) (model.Like, error)
	IsMutualLike(ctx context.Context, like model.Like) (bool, error)
}

type Dialog interface {
	CreateDialog(ctx context.Context, dialog model.Dialog) (int, error)
	GetDialogs(ctx context.Context, userId int) ([]model.Dialog, error)
	GetDialogById(ctx context.Context, id int) (model.Dialog, error)
}

type Message interface {
	GetDialogMessages(ctx context.Context, userId int, recipientId int) ([]model.Message, error)
	CreateMessage(ctx context.Context, message model.Message) (model.Message, error)
	UpdateMessage(ctx context.Context, message model.Message) (model.Message, error)
}

type Tag interface {
	GetAllTags(ctx context.Context) ([]string, error)
}

type Admin interface {
	GetAdmin(ctx context.Context, mail string) (model.Admin, error)
	CreateFeedback(ctx context.Context, stat model.Feedback) (int, error)
	CreateRecommendation(ctx context.Context, rec model.Recommendation) (int, error)
	CreateFeedFeedback(ctx context.Context, rec model.Recommendation) (int, error)
	GetFeedbacks(ctx context.Context) ([]model.Feedback, error)
	GetRecommendations(ctx context.Context) ([]model.Recommendation, error)
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
	GetComplaintTypes(ctx context.Context) ([]model.ComplaintType, error)
	CreateComplaint(ctx context.Context, complaint model.Complaint) (int, error)
	GetNextComplaint(ctx context.Context) (model.Complaint, error)
	DeleteComplaint(ctx context.Context, complaintId int) error
	AcceptComplaint(ctx context.Context, complaintId int) (model.Complaint, error)
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
