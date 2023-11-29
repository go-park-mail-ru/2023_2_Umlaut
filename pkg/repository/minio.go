package repository

import (
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioConfig struct {
	Endpoint string
	User     string
	Password string
	SSLMode  bool
}

func generatePolicy(bucketName string) string {
	return fmt.Sprintf(`{"Version": "2012-10-17","Statement": [{"Effect": "Allow","Principal": {"AWS": ["*"]},"Action": ["s3:GetBucketLocation"],"Resource": ["arn:aws:s3:::%s"]},{"Effect": "Allow","Principal": {"AWS": ["*"]},"Action": ["s3:GetObject"],"Resource": ["arn:aws:s3:::%s/*"]}]}`, bucketName, bucketName)
}

func NewMinioClient(cfg MinioConfig) (*minio.Client, error) {
	client, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.User, cfg.Password, ""),
		Secure: cfg.SSLMode,
	})
	if err != nil {
		return nil, err
	}

	return client, nil
}
