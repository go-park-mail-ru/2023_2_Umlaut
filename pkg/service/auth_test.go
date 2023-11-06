package service

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository/mocks"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAuthService_CreateUser_InvalidFields(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserRepo := mock_repository.NewMockUser(ctrl)
	mockStoreRepo := mock_repository.NewMockStore(ctrl)
	authService := NewAuthService(mockUserRepo, mockStoreRepo)
	ctx := context.Background()
	testUser := model.User{}
	mockUserRepo.EXPECT().CreateUser(ctx, gomock.Any()).Times(0)

	userID, err := authService.CreateUser(ctx, testUser)

	assert.Error(t, err)
	assert.Equal(t, 0, userID)
	assert.Equal(t, "invalid fields", err.Error())
}

func TestAuthService_GetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserRepo := mock_repository.NewMockUser(ctrl)
	mockStoreRepo := mock_repository.NewMockStore(ctrl)
	authService := NewAuthService(mockUserRepo, mockStoreRepo)
	ctx := context.Background()
	testUser := model.User{
		Mail:         "test@example.com",
		PasswordHash: "password",
	}
	mockUserRepo.EXPECT().GetUser(ctx, "test@example.com").Return(testUser, nil)

	retrievedUser, _ := authService.GetUser(ctx, "test@example.com", "password")

	assert.Equal(t, testUser, retrievedUser)
	assert.Equal(t, "password", retrievedUser.PasswordHash)
}

func TestAuthService_GetUser_InvalidPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserRepo := mock_repository.NewMockUser(ctrl)
	mockStoreRepo := mock_repository.NewMockStore(ctrl)
	authService := NewAuthService(mockUserRepo, mockStoreRepo)
	ctx := context.Background()
	testUser := model.User{}
	mockUserRepo.EXPECT().GetUser(ctx, "test@example.com").Return(testUser, nil)

	retrievedUser, err := authService.GetUser(ctx, "test@example.com", "wrong_password")

	assert.Error(t, err)
	assert.Equal(t, "invalid", err.Error())
	assert.Equal(t, model.User{}, retrievedUser)
}

func TestAuthService_GenerateCookie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserRepo := mock_repository.NewMockUser(ctrl)
	mockStoreRepo := mock_repository.NewMockStore(ctrl)
	authService := NewAuthService(mockUserRepo, mockStoreRepo)
	ctx := context.Background()
	userID := 1
	mockStoreRepo.EXPECT().SetSession(ctx, gomock.Any(), userID, 10*time.Hour).Return(nil)

	_, err := authService.GenerateCookie(ctx, userID)

	assert.NoError(t, err)
}

func TestAuthService_DeleteCookie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserRepo := mock_repository.NewMockUser(ctrl)
	mockStoreRepo := mock_repository.NewMockStore(ctrl)
	authService := NewAuthService(mockUserRepo, mockStoreRepo)
	ctx := context.Background()
	session := "session_id"
	mockStoreRepo.EXPECT().DeleteSession(ctx, "session_id").Return(nil)

	err := authService.DeleteCookie(ctx, session)

	assert.NoError(t, err)
}

func TestAuthService_GetSessionValue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserRepo := mock_repository.NewMockUser(ctrl)
	mockStoreRepo := mock_repository.NewMockStore(ctrl)
	authService := NewAuthService(mockUserRepo, mockStoreRepo)
	ctx := context.Background()
	session := "session_id"
	userID := 1
	mockStoreRepo.EXPECT().GetSession(ctx, "session_id").Return(userID, nil)

	retrievedUserID, err := authService.GetSessionValue(ctx, session)

	assert.NoError(t, err)
	assert.Equal(t, userID, retrievedUserID)
}
