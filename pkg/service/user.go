package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
)

type UserService struct {
	repoUser  repository.User
	repoStore repository.Store
	repoMinio repository.FileServer
}

func NewUserService(repoUser repository.User, repoStore repository.Store, repoMinio repository.FileServer) *UserService {
	return &UserService{repoUser: repoUser, repoStore: repoStore, repoMinio: repoMinio}
}

func (s *UserService) GetCurrentUser(ctx context.Context, userId int) (model.User, error) {
	user, err := s.repoUser.GetUserById(ctx, userId)
	if err != nil {
		return model.User{}, fmt.Errorf("GetCurrentUser error: %v", err)
	}
	user.Sanitize()

	return user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, user model.User) (model.User, error) {
	correctUser, err := s.repoUser.UpdateUser(ctx, user)
	if err != nil {
		return model.User{}, fmt.Errorf("UpdateUser error: %v", err)
	}
	correctUser.Sanitize()

	return correctUser, nil
}

func (s *UserService) UpdateUserPhoto(ctx context.Context, userId int, imagePath *string) error {
	_, err := s.repoUser.UpdateUserPhoto(ctx, userId, imagePath)
	if err != nil {
		return fmt.Errorf("UpdateUserPhoto error: %v", err)
	}
	//TODO:: сделать добавление нескольких ссылок на фото в бд
	return nil
}

func (s *UserService) CreateFile(ctx context.Context, userId int, file multipart.File, size int64) (string, error) {
	bucketName := getBucketName(userId)
	fileName := generateImageName()
	if err := s.repoMinio.CreateBucket(ctx, bucketName); err != nil {
		return "", err
	}

	buffer := make([]byte, size)
	if _, err := file.Read(buffer); err != nil {
		return "", err
	}
	if _, err := file.Seek(0, 0); err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buffer)
	err := s.repoMinio.UploadFile(ctx, bucketName, fileName, contentType, file, size)

	return fileName, err
}

func (s *UserService) GetFile(ctx context.Context, userId int, fileName string) ([]byte, string, error) {
	bucketName := getBucketName(userId)
	buffer, contentType, err := s.repoMinio.GetFile(ctx, bucketName, fileName)

	return buffer, contentType, err
}

func (s *UserService) DeleteFile(ctx context.Context, userId int, fileName string) error {
	bucketName := getBucketName(userId)
	err := s.repoMinio.DeleteFile(ctx, bucketName, fileName)

	return err
}

func generateImageName() string {
	return time.Now().String()
}

func getBucketName(userId int) string {
	return "user-id-" + strconv.Itoa(userId)
}
