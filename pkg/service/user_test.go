package service

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository/mocks"
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
			expectedError: model.InvalidUser,
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

func TestUserService_GetFile(t *testing.T) {
	mockUserId := 1
	mockFileName := "test.jpg"

	tests := []struct {
		name                string
		mockBehavior        func(r *mock_repository.MockFileServer)
		expectedBuffer      []byte
		expectedContentType string
		expectedError       error
	}{
		{
			name: "Success",
			mockBehavior: func(r *mock_repository.MockFileServer) {
				r.EXPECT().GetFile(gomock.Any(), gomock.Any(), gomock.Any()).Return([]byte("testData"), "image/jpeg", nil)
			},
			expectedBuffer:      []byte("testData"),
			expectedContentType: "image/jpeg",
			expectedError:       nil,
		},
		{
			name: "Error in GetFile",
			mockBehavior: func(r *mock_repository.MockFileServer) {
				r.EXPECT().GetFile(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, "", errors.New("some error"))
			},
			expectedBuffer:      nil,
			expectedContentType: "",
			expectedError:       errors.New("some error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockStoreRepo := mock_repository.NewMockStore(c)
			mockFileServer := mock_repository.NewMockFileServer(c)
			repoUser := mock_repository.NewMockUser(c)
			test.mockBehavior(mockFileServer)

			service := &UserService{repoUser, mockStoreRepo, mockFileServer}
			resultBuffer, resultContentType, err := service.GetFile(context.Background(), mockUserId, mockFileName)

			assert.Equal(t, test.expectedBuffer, resultBuffer)
			assert.Equal(t, test.expectedContentType, resultContentType)
			assert.Equal(t, test.expectedError, err)
		})
	}
}

func TestUserService_DeleteFile(t *testing.T) {
	mockUserId := 1
	mockFileName := "test.jpg"

	tests := []struct {
		name          string
		mockBehavior  func(r *mock_repository.MockFileServer, r1 *mock_repository.MockUser)
		expectedError error
	}{
		{
			name: "Success",
			mockBehavior: func(r *mock_repository.MockFileServer, r1 *mock_repository.MockUser) {
				s := "123"
				ps := &s
				r.EXPECT().DeleteFile(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
				r1.EXPECT().UpdateUserPhoto(gomock.Any(), mockUserId, nil).Return(ps, nil)
			},
			expectedError: nil,
		},
		{
			name: "Error in DeleteFile",
			mockBehavior: func(r *mock_repository.MockFileServer, r1 *mock_repository.MockUser) {
				r.EXPECT().DeleteFile(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("some error"))
			},
			expectedError: errors.New("DeleteFile error: some error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockStoreRepo := mock_repository.NewMockStore(c)
			mockFileServer := mock_repository.NewMockFileServer(c)
			repoUser := mock_repository.NewMockUser(c)
			test.mockBehavior(mockFileServer, repoUser)

			service := &UserService{repoUser, mockStoreRepo, mockFileServer}
			err := service.DeleteFile(context.Background(), mockUserId, mockFileName)

			assert.Equal(t, test.expectedError, err)
		})
	}
}
