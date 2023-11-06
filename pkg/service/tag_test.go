package service

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTagService_GetAllTags(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockTagRepo := mock_repository.NewMockTag(ctrl)
	tagService := NewTagService(mockTagRepo)
	testTags := []string{"Test", "test"}
	ctx := context.Background()
	mockTagRepo.EXPECT().GetAllTags(ctx).Return(testTags, nil)

	tags, err := tagService.GetAllTags(ctx)

	assert.NoError(t, err)
	assert.Equal(t, tags, testTags)
}
