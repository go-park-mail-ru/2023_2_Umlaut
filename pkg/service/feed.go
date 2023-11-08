package service

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
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

func (s *FeedService) GetNextUser(ctx context.Context, userId int) (model.User, error) {
	user, err := s.repoUser.GetUserById(ctx, userId)
	if err != nil {
		return model.User{}, fmt.Errorf("GetNextUser error: %v", err)
	}
	nextUser, err := s.repoUser.GetNextUser(ctx, user)
	if err != nil {
		return model.User{}, fmt.Errorf("GetNextUser error: %v", err)
	}

	nextUser.Sanitize()
	return nextUser, nil
}

func (s *FeedService) GetNextUsers(ctx context.Context, userId int) ([]model.User, error) {
	dialogs, err := s.repoDialog.GetDialogs(ctx, userId)
	if err != nil {
		return nil, err
	}

	var userIds []int
	for _, dialog := range dialogs {
		if dialog.User1Id == userId {
			userIds = append(userIds, dialog.User2Id)
			continue
		}
		userIds = append(userIds, dialog.User1Id)
	}

	user, err := s.repoUser.GetUserById(ctx, userId)
	if err != nil {
		return nil, err
	}

	nextUsers, err := s.repoUser.GetNextUsers(ctx, user, userIds)
	if err != nil {
		return nextUsers, err
	}
	for _, nextUser := range nextUsers {
		nextUser.Sanitize()
	}

	return nextUsers, nil
}
