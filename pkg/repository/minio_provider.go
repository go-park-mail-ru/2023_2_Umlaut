package repository

import (
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
)

type MinioProvider struct {
	client *minio.Client
}

func NewMinioProvider(client *minio.Client) *MinioProvider {
	return &MinioProvider{client: client}
}

func (m *MinioProvider) UploadFile(ctx context.Context, item model.ImageUnit) error {
	_, err := m.client.PutObject(
	   ctx,
	   bucketName,
	   item.Name,
	   item.FileLoad,
	   item.Size,
	   minio.PutObjectOptions{ContentType: item.ContentType},
	)
	//"image/png"
	return err
}

func (m *MinioProvider) DownloadFile(ctx context.Context, image string) (model.ImageUnit, error) {
	reader, err := m.client.GetObject(
	   ctx,
	   bucketName,
	   image,
	   minio.GetObjectOptions{},
	)
	if err != nil {
	   return model.ImageUnit{}, nil
	}
	defer reader.Close()
	//TODO
	return model.ImageUnit{}, nil
  }
