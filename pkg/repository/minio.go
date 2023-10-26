package repository

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioConfig struct {
	Endpoint string
	User     string
	Password string
	SSLMode  bool
}

func NewMinioClient(cfg MinioConfig) (*minio.Client , error) {
	client, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.User, cfg.Password, ""),
		Secure: cfg.SSLMode,
	})
	if err != nil {
		return nil, err
	}

	return client, nil
}
