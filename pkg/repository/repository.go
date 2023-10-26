package repository

import (
	"context"
	"database/sql"
	"io"
	"time"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
)

type User interface {
	CreateUser(ctx context.Context, user model.User) (int, error)
	GetUser(ctx context.Context, mail string) (model.User, error)
	GetUserById(ctx context.Context, id int) (model.User, error)
	GetNextUser(ctx context.Context, user model.User) (model.User, error)
	UpdateUser(ctx context.Context, user model.User) (model.User, error)
}

type Store interface {
	SetSession(ctx context.Context, SID string, id int, lifetime time.Duration) error
	GetSession(ctx context.Context, SID string) (int, error)
	DeleteSession(ctx context.Context, SID string) error
}

type FileServer interface {
	UploadFile(ctx context.Context, bucketName, fileName, contentType string, file io.Reader, size int64) error
	GetFile(ctx context.Context, bucketName, fileName string) ([]byte, string, error)
	DeleteFile(ctx context.Context, bucketName, fileName string) error
	CreateBucket(ctx context.Context, bucketName string) error
}

type Repository struct {
	User
	Store
	FileServer
}

func NewRepository(db *sql.DB, redisClient *redis.Client, minioClient *minio.Client) *Repository {
	return &Repository{
		User:  NewUserPostgres(db),
		Store: NewRedisStore(redisClient),
		FileServer: NewMinioProvider(minioClient),
	}
}
