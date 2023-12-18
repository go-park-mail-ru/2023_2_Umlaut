package service

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/constants"
	core2 "github.com/go-park-mail-ru/2023_2_Umlaut/internal/model/core"
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLikeService_CreateLike(t *testing.T) {
	mockLike := core2.Like{
		LikedByUserId: 1,
		LikedToUserId: 2,
		IsLike:        true,
	}

	tests := []struct {
		name          string
		mockBehavior  func(r *mock_repository.MockLike, d *mock_repository.MockDialog)
		expectedError error
	}{
		{
			name: "Mutual Like",
			mockBehavior: func(r *mock_repository.MockLike, d *mock_repository.MockDialog) {
				r.EXPECT().CreateLike(gomock.Any(), mockLike).Return(mockLike, nil)
				r.EXPECT().IsMutualLike(gomock.Any(), mockLike).Return(true, nil)
				d.EXPECT().CreateDialog(gomock.Any(), gomock.Any()).Return(0, nil)
				d.EXPECT().GetDialogById(gomock.Any(), gomock.Any()).Return(core2.Dialog{}, nil)
			},
			expectedError: constants.ErrMutualLike,
		},
		{
			name: "Error in GetDialogById",
			mockBehavior: func(r *mock_repository.MockLike, d *mock_repository.MockDialog) {
				r.EXPECT().CreateLike(gomock.Any(), mockLike).Return(mockLike, nil)
				r.EXPECT().IsMutualLike(gomock.Any(), mockLike).Return(true, nil)
				d.EXPECT().CreateDialog(gomock.Any(), gomock.Any()).Return(0, nil)
				d.EXPECT().GetDialogById(gomock.Any(), gomock.Any()).Return(core2.Dialog{}, errors.New("some error"))
			},
			expectedError: errors.New("some error"),
		},
		{
			name: "Non-Mutual Like",
			mockBehavior: func(r *mock_repository.MockLike, d *mock_repository.MockDialog) {
				r.EXPECT().IsMutualLike(gomock.Any(), mockLike).Return(false, nil)
				r.EXPECT().CreateLike(gomock.Any(), mockLike).Return(mockLike, nil)
			},
			expectedError: nil,
		},
		{
			name: "Error in IsMutualLike",
			mockBehavior: func(r *mock_repository.MockLike, d *mock_repository.MockDialog) {
				r.EXPECT().CreateLike(gomock.Any(), mockLike).Return(mockLike, nil)
				r.EXPECT().IsMutualLike(gomock.Any(), mockLike).Return(false, errors.New("some error"))
			},
			expectedError: errors.New("some error"),
		},
		{
			name: "Error in CreateDialog",
			mockBehavior: func(r *mock_repository.MockLike, d *mock_repository.MockDialog) {
				r.EXPECT().CreateLike(gomock.Any(), mockLike).Return(mockLike, nil)
				r.EXPECT().IsMutualLike(gomock.Any(), mockLike).Return(true, nil)
				d.EXPECT().CreateDialog(gomock.Any(), gomock.Any()).Return(0, errors.New("some error"))
			},
			expectedError: errors.New("some error"),
		},
		{
			name: "Error in CreateLike",
			mockBehavior: func(r *mock_repository.MockLike, d *mock_repository.MockDialog) {
				r.EXPECT().CreateLike(gomock.Any(), mockLike).Return(mockLike, errors.New("some error"))
			},
			expectedError: errors.New("some error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repoLike := mock_repository.NewMockLike(c)
			repoDialog := mock_repository.NewMockDialog(c)
			repoUser := mock_repository.NewMockUser(c)

			test.mockBehavior(repoLike, repoDialog)

			service := &LikeService{repoLike, repoDialog, repoUser}
			_, result := service.CreateLike(context.Background(), mockLike)

			assert.Equal(t, test.expectedError, result)
		})
	}
}
