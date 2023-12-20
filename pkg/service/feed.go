package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/constants"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/dto"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
)

type FeedService struct {
	repoUser   repository.User
	repoStore  repository.Store
	repoDialog repository.Dialog
}

func NewFeedService(repoUser repository.User, repoStore repository.Store, repoDialog repository.Dialog) *FeedService {
	return &FeedService{repoUser: repoUser, repoStore: repoStore, repoDialog: repoDialog}
}

func (s *FeedService) GetNextUser(ctx context.Context, params dto.FilterParams) (dto.FeedData, error) {
	user, err := s.repoUser.GetUserById(ctx, params.UserId)
	if errors.Is(err, constants.ErrBannedUser) {
		return dto.FeedData{}, err
	}
	if err != nil {
		return dto.FeedData{}, fmt.Errorf("GetNextUser error: %v", err)
	}
	if user.LikeCounter == 0 {
		return dto.FeedData{}, constants.ErrNoAccess
	}
	nextUser, err := s.repoUser.GetNextUser(ctx, user, params)
	if err != nil {
		return dto.FeedData{}, fmt.Errorf("GetNextUser error: %v", err)
	}

	nextUser.Sanitize()

	return dto.FeedData{User: nextUser,
		LikeCounter: user.LikeCounter,
	}, nil
}
