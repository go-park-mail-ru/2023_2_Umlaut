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
				r.EXPECT().GetNextUser(gomock.Any(), mockUser, gomock.Any()).Return(mockUser, nil)
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
				r.EXPECT().GetNextUser(gomock.Any(), mockUser, gomock.Any()).Return(model.User{}, errors.New("get next user error"))
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
			user, err := service.GetNextUser(context.Background(), model.FilterParams{})

			assert.Equal(t, test.expectedUser, user)
			assert.Equal(t, test.expectedError, err)
		})
	}
}
