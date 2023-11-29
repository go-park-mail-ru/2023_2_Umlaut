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

func TestUserService_GetCurrentUser(t *testing.T) {
	mockUserId := 1
	mockUser := model.User{
		Id:           mockUserId,
		Mail:         "max@max.ru",
		PasswordHash: "",
		Name:         "Max",
	}

	tests := []struct {
		name          string
		mockBehavior  func(r *mock_repository.MockUser)
		expectedUser  model.User
		expectedError error
	}{
		{
			name: "Success",
			mockBehavior: func(r *mock_repository.MockUser) {
				r.EXPECT().GetUserById(gomock.Any(), mockUserId).Return(mockUser, nil)
			},
			expectedUser:  mockUser,
			expectedError: nil,
		},
		{
			name: "Error in GetUserById",
			mockBehavior: func(r *mock_repository.MockUser) {
				r.EXPECT().GetUserById(gomock.Any(), mockUserId).Return(model.User{}, errors.New("some error"))
			},
			expectedUser:  model.User{},
			expectedError: errors.New("GetCurrentUser error: some error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockStoreRepo := mock_repository.NewMockStore(c)
			mockFileServer := mock_repository.NewMockFileServer(c)
			repoUser := mock_repository.NewMockUser(c)
			test.mockBehavior(repoUser)

			service := &UserService{repoUser, mockStoreRepo, mockFileServer}
			result, err := service.GetCurrentUser(context.Background(), mockUserId)

			assert.Equal(t, test.expectedUser, result)
			assert.Equal(t, test.expectedError, err)
		})
	}
}

func TestUserService_UpdateUser(t *testing.T) {
	mockUser := model.User{
		Id:           1,
		Mail:         "max@max.ru",
		PasswordHash: "passWord",
		Name:         "Max",
	}

	tests := []struct {
		name          string
		mockBehavior  func(r *mock_repository.MockUser)
		inputUser     model.User
		expectedError error
	}{
		{
			name: "Success with update password",
			mockBehavior: func(r *mock_repository.MockUser) {
				r.EXPECT().UpdateUserPassword(gomock.Any(), gomock.Any()).Return(nil)
				r.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Return(mockUser, nil)
			},
			inputUser: model.User{
				Id:           1,
				Mail:         "max@max.ru",
				PasswordHash: "passWord",
				Name:         "Max",
			},
			expectedError: nil,
		},
		{
			name:          "Invalid User",
			mockBehavior:  func(r *mock_repository.MockUser) {},
			inputUser:     model.User{PasswordHash: "passWord"},
			expectedError: static.ErrInvalidUser,
		},
		{
			name: "Error in UpdateUserPassword",
			mockBehavior: func(r *mock_repository.MockUser) {
				r.EXPECT().UpdateUserPassword(gomock.Any(), gomock.Any()).Return(errors.New("some error"))
			},
			inputUser: model.User{
				Id:           1,
				Mail:         "max@max.ru",
				PasswordHash: "passWord",
				Name:         "Max",
			},
			expectedError: errors.New("some error"),
		},
		{
			name: "Error in UpdateUser",
			mockBehavior: func(r *mock_repository.MockUser) {
				r.EXPECT().UpdateUserPassword(gomock.Any(), gomock.Any()).Return(nil)
				r.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Return(model.User{}, errors.New("some error"))
			},
			inputUser: model.User{
				Id:           1,
				Mail:         "max@max.ru",
				PasswordHash: "passWord",
				Name:         "Max",
			},
			expectedError: errors.New("some error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockStoreRepo := mock_repository.NewMockStore(c)
			mockFileServer := mock_repository.NewMockFileServer(c)
			repoUser := mock_repository.NewMockUser(c)
			test.mockBehavior(repoUser)

			service := &UserService{repoUser, mockStoreRepo, mockFileServer}
			_, err := service.UpdateUser(context.Background(), test.inputUser)

			assert.Equal(t, test.expectedError, err)
		})
	}
}
