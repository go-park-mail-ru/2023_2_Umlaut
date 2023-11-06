package service

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserService_GetCurrentUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserRepo := mock_repository.NewMockUser(ctrl)
	mockStoreRepo := mock_repository.NewMockStore(ctrl)
	mockFileServer := mock_repository.NewMockFileServer(ctrl)
	mockUserService := NewUserService(mockUserRepo, mockStoreRepo, mockFileServer)
	ctx := context.Background()
	mockUser := model.User{Id: 1}

	mockUserRepo.EXPECT().GetUserById(ctx, 1).Return(mockUser, nil)
	user, err := mockUserService.GetCurrentUser(ctx, 1)

	assert.NoError(t, err)
	assert.Equal(t, mockUser, user)
}

func TestUserService_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserRepo := mock_repository.NewMockUser(ctrl)
	mockStoreRepo := mock_repository.NewMockStore(ctrl)
	mockFileServer := mock_repository.NewMockFileServer(ctrl)
	mockUserService := NewUserService(mockUserRepo, mockStoreRepo, mockFileServer)
	ctx := context.Background()
	mockUser := model.User{Id: 1}

	mockUserRepo.EXPECT().UpdateUser(ctx, mockUser).Return(mockUser, nil)
	user, err := mockUserService.UpdateUser(ctx, mockUser)

	assert.NoError(t, err)
	assert.Equal(t, mockUser, user)
}

func TestUserService_UpdateUserPhoto(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserRepo := mock_repository.NewMockUser(ctrl)
	mockStoreRepo := mock_repository.NewMockStore(ctrl)
	mockFileServer := mock_repository.NewMockFileServer(ctrl)
	mockUserService := NewUserService(mockUserRepo, mockStoreRepo, mockFileServer)
	ctx := context.Background()
	mockPath := new(string)

	mockUserRepo.EXPECT().UpdateUserPhoto(ctx, 1, mockPath).Return(mockPath, nil)
	err := mockUserService.UpdateUserPhoto(ctx, 1, mockPath)

	assert.NoError(t, err)
}

func TestGetFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepoMinio := mock_repository.NewMockFileServer(ctrl)
	userService := NewUserService(nil, nil, mockRepoMinio)

	ctx := context.Background()
	userId := 1
	fileName := "mockFileName"
	expectedBucketName := getBucketName(userId)
	expectedContentType := "image/jpeg"
	expectedBuffer := []byte{1, 2, 3}

	mockRepoMinio.EXPECT().GetFile(ctx, expectedBucketName, fileName).Return(expectedBuffer, expectedContentType, nil)

	buffer, contentType, err := userService.GetFile(ctx, userId, fileName)

	assert.NoError(t, err)
	assert.Equal(t, expectedBuffer, buffer)
	assert.Equal(t, expectedContentType, contentType)
}

func TestDeleteFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepoMinio := mock_repository.NewMockFileServer(ctrl)
	userService := NewUserService(nil, nil, mockRepoMinio)

	ctx := context.Background()
	userId := 1
	fileName := "mockFileName"
	expectedBucketName := getBucketName(userId)

	mockRepoMinio.EXPECT().DeleteFile(ctx, expectedBucketName, fileName).Return(nil)

	err := userService.DeleteFile(ctx, userId, fileName)

	assert.NoError(t, err)
}
