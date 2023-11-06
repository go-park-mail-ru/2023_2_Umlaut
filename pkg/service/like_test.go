package service

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLikeService_CreateLike(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockLikeRepo := mock_repository.NewMockLike(ctrl)
	mockLikeService := NewLikeService(mockLikeRepo)
	ctx := context.Background()
	mockLike := model.Like{}
	mockLikeRepo.EXPECT().CreateLike(ctx, mockLike).Return(mockLike, nil)

	err := mockLikeService.CreateLike(ctx, mockLike)

	assert.NoError(t, err)
}

func TestLikeService_IsUserLiked(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLikeRepo := mock_repository.NewMockLike(ctrl)
	mockLikeService := NewLikeService(mockLikeRepo)
	ctx := context.Background()
	mockLike := model.Like{}
	mockLikeRepo.EXPECT().Exists(ctx, mockLike).Return(true, nil)

	exist, err := mockLikeService.IsUserLiked(ctx, mockLike)

	assert.NoError(t, err)
	assert.Equal(t, exist, true)
}

func TestLikeService_IsLikeExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockLikeRepo := mock_repository.NewMockLike(ctrl)
	mockLikeService := NewLikeService(mockLikeRepo)
	ctx := context.Background()
	mockLike := model.Like{}
	mockLikeRepo.EXPECT().Exists(ctx, mockLike).Return(true, nil)

	exist, err := mockLikeService.IsLikeExists(ctx, mockLike)

	assert.NoError(t, err)
	assert.Equal(t, exist, true)
}
