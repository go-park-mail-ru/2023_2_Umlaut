package service

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository/mocks"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAuthService_CreateUser(t *testing.T) {
	mockUser := model.User{
		Mail:         "test@example.com",
		PasswordHash: "password",
		Name:         "TestUser",
	}

	tests := []struct {
		name          string
		mockBehavior  func(r *mock_repository.MockUser)
		expectedID    int
		expectedError error
	}{
		{
			name: "Success",
			mockBehavior: func(r *mock_repository.MockUser) {
				r.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(1, nil)
			},
			expectedID:    1,
			expectedError: nil,
		},
		{
			name: "Invalid Fields",
			mockBehavior: func(r *mock_repository.MockUser) {
				r.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(0, errors.New("invalid fields"))
			},
			expectedID:    0,
			expectedError: errors.New("invalid fields"),
		},
		{
			name: "Already Exists",
			mockBehavior: func(r *mock_repository.MockUser) {
				r.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(0, model.AlreadyExists)
			},
			expectedID:    0,
			expectedError: model.AlreadyExists,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repoUser := mock_repository.NewMockUser(c)
			test.mockBehavior(repoUser)

			service := &AuthService{repoUser: repoUser}
			id, err := service.CreateUser(context.Background(), mockUser)

			assert.Equal(t, test.expectedID, id)
			assert.Equal(t, test.expectedError, err)
		})
	}
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
