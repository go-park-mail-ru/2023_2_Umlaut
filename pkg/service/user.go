package service

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
)

type UserService struct {
	repoUser  repository.User
	repoStore repository.Store
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
