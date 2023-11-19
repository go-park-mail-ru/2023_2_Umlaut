package service

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository/mocks"
	"github.com/go-park-mail-ru/2023_2_Umlaut/static"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
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
				r.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(0, static.ErrAlreadyExists)
			},
			expectedID:    0,
			expectedError: static.ErrAlreadyExists,
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
	mockUser := model.User{
		Mail:         "test@example.com",
		PasswordHash: "",
		Name:         "TestUser",
		Salt:         "",
	}

	tests := []struct {
		name          string
		mockBehavior  func(r *mock_repository.MockUser)
		expectedUser  model.User
		expectedError error
	}{
		{
			name: "User Not Found",
			mockBehavior: func(r *mock_repository.MockUser) {
				r.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(model.User{}, errors.New("user not found"))
			},
			expectedUser:  model.User{},
			expectedError: errors.New("user not found"),
		},
		{
			name: "Invalid Password",
			mockBehavior: func(r *mock_repository.MockUser) {
				r.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(mockUser, nil)
			},
			expectedUser:  mockUser,
			expectedError: errors.New("invalid"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repoUser := mock_repository.NewMockUser(c)
			test.mockBehavior(repoUser)

			service := &AuthService{repoUser: repoUser}
			user, err := service.GetUser(context.Background(), "test@example.com", "password")

			assert.Equal(t, test.expectedUser, user)
			assert.Equal(t, test.expectedError, err)
		})
	}
}

func TestAuthService_GenerateCookie(t *testing.T) {
	tests := []struct {
		name          string
		mockBehavior  func(r *mock_repository.MockStore)
		expectedError error
	}{
		{
			name: "Success",
			mockBehavior: func(r *mock_repository.MockStore) {
				r.EXPECT().SetSession(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Error Setting Session",
			mockBehavior: func(r *mock_repository.MockStore) {
				r.EXPECT().SetSession(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("session error"))
			},
			expectedError: errors.New("session error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repoStore := mock_repository.NewMockStore(c)
			test.mockBehavior(repoStore)

			service := &AuthService{repoStore: repoStore}
			_, err := service.GenerateCookie(context.Background(), 1)

			assert.Equal(t, test.expectedError, err)
		})
	}
}

func TestAuthService_DeleteCookie(t *testing.T) {
	tests := []struct {
		name          string
		mockBehavior  func(r *mock_repository.MockStore)
		expectedError error
	}{
		{
			name: "Success",
			mockBehavior: func(r *mock_repository.MockStore) {
				r.EXPECT().DeleteSession(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Error Deleting Session",
			mockBehavior: func(r *mock_repository.MockStore) {
				r.EXPECT().DeleteSession(gomock.Any(), gomock.Any()).Return(errors.New("delete session error"))
			},
			expectedError: errors.New("delete session error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repoStore := mock_repository.NewMockStore(c)
			test.mockBehavior(repoStore)

			service := &AuthService{repoStore: repoStore}
			err := service.DeleteCookie(context.Background(), "sessionID")

			assert.Equal(t, test.expectedError, err)
		})
	}
}

func TestAuthService_GetSessionValue(t *testing.T) {
	tests := []struct {
		name          string
		mockBehavior  func(r *mock_repository.MockStore)
		expectedID    int
		expectedError error
	}{
		{
			name: "Success",
			mockBehavior: func(r *mock_repository.MockStore) {
				r.EXPECT().GetSession(gomock.Any(), gomock.Any()).Return(1, nil)
			},
			expectedID:    1,
			expectedError: nil,
		},
		{
			name: "Error Getting Session",
			mockBehavior: func(r *mock_repository.MockStore) {
				r.EXPECT().GetSession(gomock.Any(), gomock.Any()).Return(0, errors.New("get session error"))
			},
			expectedID:    0,
			expectedError: errors.New("get session error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repoStore := mock_repository.NewMockStore(c)
			test.mockBehavior(repoStore)

			service := &AuthService{repoStore: repoStore}
			id, err := service.GetSessionValue(context.Background(), "sessionID")

			assert.Equal(t, test.expectedID, id)
			assert.Equal(t, test.expectedError, err)
		})
	}
}
