package service

import (
	"context"
	"errors"
	"testing"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	mock_repository "github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository/mocks"
	"github.com/go-park-mail-ru/2023_2_Umlaut/static"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestFeedService_GetNextUser(t *testing.T) {
	mockUser := model.User{
		Id:   1,
		Name: "TestUser",
	}
	mockFeedData := model.FeedData{
		User: mockUser,
	}

	tests := []struct {
		name          string
		mockBehavior  func(r *mock_repository.MockUser)
		expectedFeed  model.FeedData
		expectedError error
	}{
		{
			name: "Success",
			mockBehavior: func(r *mock_repository.MockUser) {
				r.EXPECT().GetUserById(gomock.Any(), 1).Return(mockUser, nil)
				r.EXPECT().GetNextUser(gomock.Any(), mockUser, gomock.Any()).Return(mockUser, nil)
			},
			expectedFeed:  mockFeedData,
			expectedError: nil,
		},
		{
			name: "Error Getting User",
			mockBehavior: func(r *mock_repository.MockUser) {
				r.EXPECT().GetUserById(gomock.Any(), 1).Return(model.User{}, errors.New("get user error"))
			},
			expectedFeed:  model.FeedData{},
			expectedError: errors.New("GetNextUser error: get user error"),
		},
		{
			name: "Banned User",
			mockBehavior: func(r *mock_repository.MockUser) {
				r.EXPECT().GetUserById(gomock.Any(), 1).Return(model.User{}, static.ErrBannedUser)
			},
			expectedFeed:  model.FeedData{},
			expectedError: static.ErrBannedUser,
		},
		{
			name: "Error Getting Next User",
			mockBehavior: func(r *mock_repository.MockUser) {
				r.EXPECT().GetUserById(gomock.Any(), 1).Return(mockUser, nil)
				r.EXPECT().GetNextUser(gomock.Any(), mockUser, gomock.Any()).Return(model.User{}, errors.New("get next user error"))
			},
			expectedFeed:  model.FeedData{},
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
			feedData, err := service.GetNextUser(context.Background(), model.FilterParams{UserId: 1})

			assert.Equal(t, test.expectedFeed, feedData)
			assert.Equal(t, test.expectedError, err)
		})
	}
}
