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

func (s *FeedService) GetNextUser(ctx context.Context, session string) (model.User, error) {
	id, err := s.repoStore.GetSession(ctx, session)
	if err != nil {
		return model.User{}, err
	}

	user, _ := s.repoUser.GetUserById(id)
	nextUser, err := s.repoUser.GetNextUser(user)
	if err != nil {
		return model.User{}, err
	}
	nextUser.Sanitize()

	return nextUser, nil
}
