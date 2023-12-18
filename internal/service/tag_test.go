package service

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTagService_GetAllTags(t *testing.T) {
	mockTags := []string{"tag1", "tag2", "tag3"}

	tests := []struct {
		name          string
		mockBehavior  func(r *mock_repository.MockTag)
		expectedList  []string
		expectedError error
	}{
		{
			name: "Success",
			mockBehavior: func(r *mock_repository.MockTag) {
				r.EXPECT().GetAllTags(gomock.Any()).Return(mockTags, nil)
			},
			expectedList:  mockTags,
			expectedError: nil,
		},
		{
			name: "Error Getting Tags",
			mockBehavior: func(r *mock_repository.MockTag) {
				r.EXPECT().GetAllTags(gomock.Any()).Return(nil, errors.New("get tags error"))
			},
			expectedList:  nil,
			expectedError: errors.New("get tags error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repoTag := mock_repository.NewMockTag(c)
			test.mockBehavior(repoTag)

			service := &TagService{repoTag: repoTag}
			tags, err := service.GetAllTags(context.Background())

			assert.Equal(t, test.expectedList, tags)
			assert.Equal(t, test.expectedError, err)
		})
	}
}
