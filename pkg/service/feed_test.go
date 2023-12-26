package service

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/constants"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/core"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/dto"
	"testing"

	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestFeedService_GetNextUser(t *testing.T) {
	mockUser := core.User{
		Id:   1,
		Name: "TestUser",
		LikeCounter: 50,
	}
	mockFeedData := dto.FeedData{
		User: mockUser,
		LikeCounter: 50,
	}

	tests := []struct {
		name          string
		mockBehavior  func(r *mock_repository.MockUser)
		expectedFeed  dto.FeedData
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
				r.EXPECT().GetUserById(gomock.Any(), 1).Return(core.User{}, errors.New("get user error"))
			},
			expectedFeed:  dto.FeedData{},
			expectedError: errors.New("GetNextUser error: get user error"),
		},
		{
			name: "Banned User",
			mockBehavior: func(r *mock_repository.MockUser) {
				r.EXPECT().GetUserById(gomock.Any(), 1).Return(core.User{}, constants.ErrBannedUser)
			},
			expectedFeed:  dto.FeedData{},
			expectedError: constants.ErrBannedUser,
		},
		{
			name: "Error Getting Next User",
			mockBehavior: func(r *mock_repository.MockUser) {
				r.EXPECT().GetUserById(gomock.Any(), 1).Return(mockUser, nil)
				r.EXPECT().GetNextUser(gomock.Any(), mockUser, gomock.Any()).Return(core.User{}, errors.New("get next user error"))
			},
			expectedFeed:  dto.FeedData{},
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
			feedData, err := service.GetNextUser(context.Background(), dto.FilterParams{UserId: 1})

			assert.Equal(t, test.expectedFeed, feedData)
			assert.Equal(t, test.expectedError, err)
		})
	}
}
