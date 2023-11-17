package repository

import (
	"context"
	"io"

	"fmt"

	"github.com/minio/minio-go/v7"
)

type MinioProvider struct {
	client *minio.Client
}

func NewMinioProvider(client *minio.Client) *MinioProvider {
	return &MinioProvider{client: client}
}

func (m *MinioProvider) UploadFile(ctx context.Context, bucketName, fileName, contentType string, file io.Reader, size int64) error {
	_, err := m.client.PutObject(
		ctx,
		"user-id-1",
		fileName,
		file,
		size,
		minio.PutObjectOptions{ContentType: contentType},
	)
	if err != nil {
		return fmt.Errorf("failed to upload file. err: %w", err)
	}
	return nil
}

func (m *MinioProvider) DeleteFile(ctx context.Context, bucketName, fileName string) error {
	err := m.client.RemoveObject(ctx, "user-id-1", fileName, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete file. err: %w", err)
	}

	return nil
}

func (m *MinioProvider) CreateBucket(ctx context.Context, bucketName string) error {
	exists, err := m.client.BucketExists(ctx, bucketName)
	if err != nil {
		return fmt.Errorf("failed to create bucket. err: %w", err)
	}
	if exists {
		return nil
	}
	err = m.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		return fmt.Errorf("failed to create bucket. err: %w", err)
	}
	err = m.client.SetBucketPolicy(ctx, bucketName, generatePolicy(bucketName))
	return err
}
