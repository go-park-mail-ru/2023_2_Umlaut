package service

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFeedService_GetNextUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserRepo := mock_repository.NewMockUser(ctrl)
	feedService := NewFeedService(mockUserRepo, nil, nil)
	ctx := context.Background()
	userId := 1
	expectedUser := model.User{
		Id:   userId,
		Name: "John Doe",
	}
	mockUserRepo.EXPECT().GetUserById(ctx, userId).Return(expectedUser, nil)
	expectedNextUser := model.User{
		Id:   userId + 1,
		Name: "Next User",
	}
	mockUserRepo.EXPECT().GetNextUser(ctx, expectedUser).Return(expectedNextUser, nil)

	nextUser, err := feedService.GetNextUser(ctx, userId)

	assert.NoError(t, err)
	assert.Equal(t, expectedNextUser, nextUser)
}

func TestFeedService_GetNextUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDialogRepo := mock_repository.NewMockDialog(ctrl)
	mockUserRepo := mock_repository.NewMockUser(ctrl)
	feedService := NewFeedService(mockUserRepo, nil, mockDialogRepo)
	ctx := context.Background()
	userId := 1
	expectedDialogs := []model.Dialog{
		{Id: 1, User1Id: userId, User2Id: 2},
		{Id: 2, User1Id: userId, User2Id: 3},
	}
	mockDialogRepo.EXPECT().GetDialogs(ctx, userId).Return(expectedDialogs, nil)
	expectedUser := model.User{
		Id:   userId,
		Name: "John Doe",
	}
	mockUserRepo.EXPECT().GetUserById(ctx, userId).Return(expectedUser, nil)
	userIds := []int{2, 3}
	expectedNextUsers := []model.User{
		{Id: 2, Name: "User 2"},
		{Id: 3, Name: "User 3"},
	}
	mockUserRepo.EXPECT().GetNextUsers(ctx, expectedUser, userIds).Return(expectedNextUsers, nil)

	nextUsers, err := feedService.GetNextUsers(ctx, userId)

	assert.NoError(t, err)
	expectedUserNames := []string{"User 2", "User 3"}
	for i, user := range nextUsers {
		assert.Equal(t, expectedUserNames[i], user.Name)
	}
}
