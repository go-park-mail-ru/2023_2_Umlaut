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
		bucketName,
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

func (m *MinioProvider) GetFile(ctx context.Context, bucketName, fileName string) ([]byte, string, error) {
	obj, err := m.client.GetObject(
		ctx,
		bucketName,
		fileName,
		minio.GetObjectOptions{},
	)
	if err != nil {
		return nil, "", fmt.Errorf("failed to get file. err: %w", err)
	}
	defer obj.Close()

	objectInfo, err := obj.Stat()
	if err != nil {
		return nil, "", fmt.Errorf("failed to get file. err: %w", err)
	}
	buffer := make([]byte, objectInfo.Size)
	_, err = obj.Read(buffer)
	if err != nil && err != io.EOF {
		return nil, "", fmt.Errorf("failed to get file. err: %w", err)
	}

	return buffer, objectInfo.ContentType, nil
}

func (m *MinioProvider) DeleteFile(ctx context.Context, bucketName, fileName string) error {
	err := m.client.RemoveObject(ctx, bucketName, fileName, minio.RemoveObjectOptions{})
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
	return nil
}
