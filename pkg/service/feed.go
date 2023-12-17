package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
	"github.com/go-park-mail-ru/2023_2_Umlaut/static"
)

type FeedService struct {
	repoUser   repository.User
	repoStore  repository.Store
	repoDialog repository.Dialog
}

func NewFeedService(repoUser repository.User, repoStore repository.Store, repoDialog repository.Dialog) *FeedService {
	return &FeedService{repoUser: repoUser, repoStore: repoStore, repoDialog: repoDialog}
}

func (s *FeedService) GetNextUser(ctx context.Context, params model.FilterParams) (model.FeedData, error) {
	user, err := s.repoUser.GetUserById(ctx, params.UserId)
	if errors.Is(err, static.ErrBannedUser) {
		return model.FeedData{}, err
	}
	if err != nil {
		return model.FeedData{}, fmt.Errorf("GetNextUser error: %v", err)
	}
	nextUser, err := s.repoUser.GetNextUser(ctx, user, params)
	if err != nil {
		return model.FeedData{}, fmt.Errorf("GetNextUser error: %v", err)
	}
	
	nextUser.Sanitize()

	feed := model.FeedData{User: user, LikeCounter: nextUser.LikeCounter}
	return feed, nil
}
