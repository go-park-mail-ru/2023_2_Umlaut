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

func TestFeedService_GetNextUser(t *testing.T) {
	mockUser := model.User{
		Id:   1,
		Name: "TestUser",
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
				r.EXPECT().GetUserById(gomock.Any(), 1).Return(mockUser, nil)
				r.EXPECT().GetNextUser(gomock.Any(), mockUser).Return(mockUser, nil)
			},
			expectedUser:  mockUser,
			expectedError: nil,
		},
		{
			name: "Error Getting User",
			mockBehavior: func(r *mock_repository.MockUser) {
				r.EXPECT().GetUserById(gomock.Any(), 1).Return(model.User{}, errors.New("get user error"))
			},
			expectedUser:  model.User{},
			expectedError: errors.New("GetNextUser error: get user error"),
		},
		{
			name: "Error Getting Next User",
			mockBehavior: func(r *mock_repository.MockUser) {
				r.EXPECT().GetUserById(gomock.Any(), 1).Return(mockUser, nil)
				r.EXPECT().GetNextUser(gomock.Any(), mockUser).Return(model.User{}, errors.New("get next user error"))
			},
			expectedUser:  model.User{},
			expectedError: errors.New("GetNextUser error: get next user error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repoUser := mock_repository.NewMockUser(c)
			test.mockBehavior(repoUser)

			service := &FeedService{repoUser: repoUser}
			user, err := service.GetNextUser(context.Background(), 1)

			assert.Equal(t, test.expectedUser, user)
			assert.Equal(t, test.expectedError, err)
		})
	}
}

func TestFeedService_GetNextUsers(t *testing.T) {
	mockDialogs := []model.Dialog{
		{Id: 1, User1Id: 1, User2Id: 2},
		{Id: 2, User1Id: 1, User2Id: 3},
	}

	mockUsers := []model.User{
		{Id: 2, Name: "User2"},
		{Id: 3, Name: "User3"},
	}

	tests := []struct {
		name          string
		mockBehavior  func(r *mock_repository.MockUser, rDialog *mock_repository.MockDialog)
		expectedList  []model.User
		expectedError error
	}{
		{
			name: "Success",
			mockBehavior: func(r *mock_repository.MockUser, rDialog *mock_repository.MockDialog) {
				rDialog.EXPECT().GetDialogs(gomock.Any(), 1).Return(mockDialogs, nil)
				r.EXPECT().GetUserById(gomock.Any(), 1).Return(mockUsers[0], nil)
				r.EXPECT().GetNextUsers(gomock.Any(), mockUsers[0], []int{2, 3}).Return(mockUsers, nil)
			},
			expectedList:  mockUsers,
			expectedError: nil,
		},
		{
			name: "Error Getting Dialogs",
			mockBehavior: func(r *mock_repository.MockUser, rDialog *mock_repository.MockDialog) {
				rDialog.EXPECT().GetDialogs(gomock.Any(), 1).Return(nil, errors.New("get dialogs error"))
			},
			expectedList:  nil,
			expectedError: errors.New("get dialogs error"),
		},
		{
			name: "Error Getting User",
			mockBehavior: func(r *mock_repository.MockUser, rDialog *mock_repository.MockDialog) {
				rDialog.EXPECT().GetDialogs(gomock.Any(), 1).Return(mockDialogs, nil)
				r.EXPECT().GetUserById(gomock.Any(), 1).Return(model.User{}, errors.New("get user error"))
			},
			expectedList:  nil,
			expectedError: errors.New("get user error"),
		},
		{
			name: "Error Getting Next Users",
			mockBehavior: func(r *mock_repository.MockUser, rDialog *mock_repository.MockDialog) {
				rDialog.EXPECT().GetDialogs(gomock.Any(), 1).Return(mockDialogs, nil)
				r.EXPECT().GetUserById(gomock.Any(), 1).Return(mockUsers[0], nil)
				r.EXPECT().GetNextUsers(gomock.Any(), mockUsers[0], []int{2, 3}).Return(nil, errors.New("get next users error"))
			},
			expectedList:  nil,
			expectedError: errors.New("get next users error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repoUser := mock_repository.NewMockUser(c)
			repoDialog := mock_repository.NewMockDialog(c)
			test.mockBehavior(repoUser, repoDialog)

			service := &FeedService{repoUser: repoUser, repoDialog: repoDialog}
			users, err := service.GetNextUsers(context.Background(), 1)

			assert.Equal(t, test.expectedList, users)
			assert.Equal(t, test.expectedError, err)
		})
	}
}
