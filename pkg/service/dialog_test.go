package service

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDialogService_CreateDialog(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDialogRepo := mock_repository.NewMockDialog(ctrl)
	dialogService := NewDialogService(mockDialogRepo)
	ctx := context.Background()
	testDialog := model.Dialog{}
	mockDialogRepo.EXPECT().CreateDialog(ctx, testDialog).Return(1, nil)

	id, err := dialogService.CreateDialog(ctx, testDialog)

	assert.NoError(t, err)
	assert.Equal(t, 1, id)
}

func TestDialogService_GetDialogs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDialogRepo := mock_repository.NewMockDialog(ctrl)
	dialogService := NewDialogService(mockDialogRepo)
	ctx := context.Background()
	var testDialogs []model.Dialog
	mockDialogRepo.EXPECT().GetDialogs(ctx, 1).Return(testDialogs, nil)

	dialogs, err := dialogService.GetDialogs(ctx, 1)

	assert.NoError(t, err)
	assert.Equal(t, dialogs, testDialogs)
}
