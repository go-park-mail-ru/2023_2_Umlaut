package service

import (
	"context"
	"time"
	"strconv"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
)

type UserService struct {
	repoUser  repository.User
	repoStore repository.Store
	repoMinio repository.MinioProvider
}

func NewUserService(repoUser repository.User, repoStore repository.Store) *UserService {
	return &UserService{repoUser: repoUser, repoStore: repoStore}
}

func (s *UserService) GetCurrentUser(ctx context.Context, userId int) (model.User, error) {
	user, err := s.repoUser.GetUserById(ctx, userId)
	if err != nil {
		return model.User{}, err
	}
	user.Sanitize()

	return user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, user model.User) (model.User, error) {
	correctUser, err := s.repoUser.UpdateUser(ctx, user)
	if err != nil {
		return model.User{}, err
	}
	correctUser.Sanitize()

	return correctUser, nil
}

func (s *UserService) UpdatePhoto(ctx context.Context, userId int, img model.ImageUnit) (string, error) {
	img.Name = generateImageName(userId)
	err := s.repoMinio.UploadFile(ctx, img)

	return img.Name, err
}

func generateImageName(userId int) string {
	return strconv.Itoa(userId) + "/" + time.Now().String()
}
