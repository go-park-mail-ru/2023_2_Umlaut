package service

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
)

type FeedService struct {
	repoUser  repository.User
	repoStore repository.Store
}

func NewFeedService(repoUser repository.User, repoStore repository.Store) *FeedService {
	return &FeedService{repoUser: repoUser, repoStore: repoStore}
}

func (s *FeedService) GetNextUser(ctx context.Context, userId int) (model.User, error) {
	user, _ := s.repoUser.GetUserById(ctx, userId)
	nextUser, err := s.repoUser.GetNextUser(ctx, user)
	if err != nil {
		return model.User{}, err
	}

	return nextUser, nil
}
