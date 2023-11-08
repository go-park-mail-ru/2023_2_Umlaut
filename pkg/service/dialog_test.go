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

func TestDialogService_CreateDialog(t *testing.T) {
	mockDialog := model.Dialog{
		Id:      1,
		User1Id: 1,
		User2Id: 2,
	}

	tests := []struct {
		name          string
		mockBehavior  func(r *mock_repository.MockDialog)
		expectedID    int
		expectedError error
	}{
		{
			name: "Success",
			mockBehavior: func(r *mock_repository.MockDialog) {
				r.EXPECT().CreateDialog(gomock.Any(), mockDialog).Return(1, nil)
			},
			expectedID:    1,
			expectedError: nil,
		},
		{
			name: "Error Creating Dialog",
			mockBehavior: func(r *mock_repository.MockDialog) {
				r.EXPECT().CreateDialog(gomock.Any(), mockDialog).Return(0, errors.New("create dialog error"))
			},
			expectedID:    0,
			expectedError: errors.New("create dialog error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repoDialog := mock_repository.NewMockDialog(c)
			test.mockBehavior(repoDialog)

			service := &DialogService{repoDialog: repoDialog}
			id, err := service.CreateDialog(context.Background(), mockDialog)

			assert.Equal(t, test.expectedID, id)
			assert.Equal(t, test.expectedError, err)
		})
	}
}

func TestDialogService_GetDialogs(t *testing.T) {
	mockDialogs := []model.Dialog{
		{Id: 1, User1Id: 1, User2Id: 2},
		{Id: 2, User1Id: 3, User2Id: 4},
	}

	tests := []struct {
		name          string
		mockBehavior  func(r *mock_repository.MockDialog)
		expectedList  []model.Dialog
		expectedError error
	}{
		{
			name: "Success",
			mockBehavior: func(r *mock_repository.MockDialog) {
				r.EXPECT().GetDialogs(gomock.Any(), 1).Return(mockDialogs, nil)
			},
			expectedList:  mockDialogs,
			expectedError: nil,
		},
		{
			name: "Error Getting Dialogs",
			mockBehavior: func(r *mock_repository.MockDialog) {
				r.EXPECT().GetDialogs(gomock.Any(), 1).Return(nil, errors.New("get dialogs error"))
			},
			expectedList:  nil,
			expectedError: errors.New("get dialogs error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repoDialog := mock_repository.NewMockDialog(c)
			test.mockBehavior(repoDialog)

			service := &DialogService{repoDialog: repoDialog}
			dialogs, err := service.GetDialogs(context.Background(), 1)

			assert.Equal(t, test.expectedList, dialogs)
			assert.Equal(t, test.expectedError, err)
		})
	}
}
