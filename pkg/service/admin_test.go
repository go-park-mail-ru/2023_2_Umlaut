package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/core"
	"testing"

	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAdminService_CreateRecommendation(t *testing.T) {
	mockID := 1

	tests := []struct {
		name           string
		inputRec       core.Recommendation
		mockBehavior   func(r *mock_repository.MockAdmin)
		expectedResult int
		expectedError  error
	}{
		{
			name:     "Success",
			inputRec: core.Recommendation{},
			mockBehavior: func(r *mock_repository.MockAdmin) {
				r.EXPECT().CreateRecommendation(gomock.Any(), gomock.Any()).Return(mockID, nil)
			},
			expectedResult: mockID,
			expectedError:  nil,
		},
		{
			name:     "Error",
			inputRec: core.Recommendation{},
			mockBehavior: func(r *mock_repository.MockAdmin) {
				r.EXPECT().CreateRecommendation(gomock.Any(), gomock.Any()).Return(0, errors.New("error creating recommendation"))
			},
			expectedResult: 0,
			expectedError:  errors.New("error creating recommendation"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repoAdmin := mock_repository.NewMockAdmin(ctrl)
			test.mockBehavior(repoAdmin)

			service := &AdminService{RepoAdmin: repoAdmin}
			result, err := service.CreateRecommendation(context.Background(), test.inputRec)

			assert.Equal(t, test.expectedResult, result)
			assert.Equal(t, test.expectedError, err)
		})
	}
}

func TestAdminService_CreateFeedFeedback(t *testing.T) {
	mockID := 1

	tests := []struct {
		name           string
		inputRec       core.Recommendation
		mockBehavior   func(r *mock_repository.MockAdmin)
		expectedResult int
		expectedError  error
	}{
		{
			name:     "Success",
			inputRec: core.Recommendation{},
			mockBehavior: func(r *mock_repository.MockAdmin) {
				r.EXPECT().CreateFeedFeedback(gomock.Any(), gomock.Any()).Return(mockID, nil)
			},
			expectedResult: mockID,
			expectedError:  nil,
		},
		{
			name:     "Error",
			inputRec: core.Recommendation{},
			mockBehavior: func(r *mock_repository.MockAdmin) {
				r.EXPECT().CreateFeedFeedback(gomock.Any(), gomock.Any()).Return(0, errors.New("error creating feed feedback"))
			},
			expectedResult: 0,
			expectedError:  errors.New("error creating feed feedback"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repoAdmin := mock_repository.NewMockAdmin(ctrl)
			test.mockBehavior(repoAdmin)

			service := &AdminService{RepoAdmin: repoAdmin}
			result, err := service.CreateFeedFeedback(context.Background(), test.inputRec)

			assert.Equal(t, test.expectedResult, result)
			assert.Equal(t, test.expectedError, err)
		})
	}
}

func TestAdminService_CreateFeedback(t *testing.T) {
	mockID := 1

	tests := []struct {
		name           string
		inputStat      core.Feedback
		mockBehavior   func(r *mock_repository.MockAdmin)
		expectedResult int
		expectedError  error
	}{
		{
			name:      "Success",
			inputStat: core.Feedback{},
			mockBehavior: func(r *mock_repository.MockAdmin) {
				r.EXPECT().CreateFeedback(gomock.Any(), gomock.Any()).Return(mockID, nil)
			},
			expectedResult: mockID,
			expectedError:  nil,
		},
		{
			name:      "Error",
			inputStat: core.Feedback{},
			mockBehavior: func(r *mock_repository.MockAdmin) {
				r.EXPECT().CreateFeedback(gomock.Any(), gomock.Any()).Return(0, errors.New("error creating feedback"))
			},
			expectedResult: 0,
			expectedError:  errors.New("error creating feedback"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repoAdmin := mock_repository.NewMockAdmin(ctrl)
			test.mockBehavior(repoAdmin)

			service := &AdminService{RepoAdmin: repoAdmin}
			result, err := service.CreateFeedback(context.Background(), test.inputStat)

			assert.Equal(t, test.expectedResult, result)
			assert.Equal(t, test.expectedError, err)
		})
	}
}

func TestAdminService_GetCSATType(t *testing.T) {
	tests := []struct {
		name               string
		userID             int
		mockBehavior       func(r *mock_repository.MockAdmin, u *mock_repository.MockUser)
		expectedResult     int
		expectedError      error
		expectedErrorCause string
	}{
		{
			name:   "CSAT type 0",
			userID: 1,
			mockBehavior: func(r *mock_repository.MockAdmin, u *mock_repository.MockUser) {
				u.EXPECT().ShowCSAT(gomock.Any(), 1).Return(true, nil)
				r.EXPECT().ShowFeedback(gomock.Any(), 1).Return(false, nil)
				r.EXPECT().ShowRecommendation(gomock.Any(), 1).Return(false, nil)
			},
			expectedResult: 0,
			expectedError:  nil,
		},
		{
			name:   "CSAT type 1",
			userID: 1,
			mockBehavior: func(r *mock_repository.MockAdmin, u *mock_repository.MockUser) {
				u.EXPECT().ShowCSAT(gomock.Any(), 1).Return(true, nil)
				r.EXPECT().ShowFeedback(gomock.Any(), 1).Return(true, nil)
			},
			expectedResult: 1,
			expectedError:  nil,
		},
		{
			name:   "CSAT type 2",
			userID: 2,
			mockBehavior: func(r *mock_repository.MockAdmin, u *mock_repository.MockUser) {
				u.EXPECT().ShowCSAT(gomock.Any(), 2).Return(true, nil)
				r.EXPECT().ShowFeedback(gomock.Any(), 2).Return(false, nil)
				r.EXPECT().ShowRecommendation(gomock.Any(), 2).Return(true, nil)
			},
			expectedResult: 2,
			expectedError:  nil,
		},
		{
			name:   "No CSAT",
			userID: 3,
			mockBehavior: func(r *mock_repository.MockAdmin, u *mock_repository.MockUser) {
				u.EXPECT().ShowCSAT(gomock.Any(), 3).Return(false, nil)
			},
			expectedResult: 0,
			expectedError:  nil,
		},
		{
			name:   "Error ShowCSAT",
			userID: 4,
			mockBehavior: func(r *mock_repository.MockAdmin, u *mock_repository.MockUser) {
				u.EXPECT().ShowCSAT(gomock.Any(), 4).Return(false, errors.New("error getting CSAT"))
			},
			expectedResult: 0,
			expectedError:  fmt.Errorf("GetCSATType error: error getting CSAT"),
		},
		{
			name:   "Error ShowFeedback",
			userID: 5,
			mockBehavior: func(r *mock_repository.MockAdmin, u *mock_repository.MockUser) {
				u.EXPECT().ShowCSAT(gomock.Any(), 5).Return(true, nil)
				r.EXPECT().ShowFeedback(gomock.Any(), 5).Return(false, errors.New("error ShowFeedback"))
			},
			expectedResult: 0,
			expectedError:  fmt.Errorf("GetCSATType error: error ShowFeedback"),
		},
		{
			name:   "Error ShowRecommendation",
			userID: 6,
			mockBehavior: func(r *mock_repository.MockAdmin, u *mock_repository.MockUser) {
				u.EXPECT().ShowCSAT(gomock.Any(), 6).Return(true, nil)
				r.EXPECT().ShowFeedback(gomock.Any(), 6).Return(false, nil)
				r.EXPECT().ShowRecommendation(gomock.Any(), 6).Return(true, errors.New("error ShowRecommendation"))
			},
			expectedResult: 0,
			expectedError:  fmt.Errorf("GetCSATType error: error ShowRecommendation"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repoAdmin := mock_repository.NewMockAdmin(ctrl)
			repoUser := mock_repository.NewMockUser(ctrl)
			test.mockBehavior(repoAdmin, repoUser)

			service := &AdminService{RepoAdmin: repoAdmin, RepoUser: repoUser}
			result, err := service.GetCSATType(context.Background(), test.userID)

			assert.Equal(t, test.expectedResult, result)
			assert.Equal(t, test.expectedError, err)
		})
	}
}
