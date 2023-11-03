package service

import (
	"context"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
)

type DialogService struct {
	repoDialog repository.Dialog
}

func NewDialogService(repoDialog repository.Dialog) *DialogService {
	return &DialogService{repoDialog: repoDialog}
}

func (s *DialogService) CreateDialog(ctx context.Context, dialog model.Dialog) (int, error) {
	if dialog.User1Id > dialog.User2Id {
		tmp := dialog.User1Id
		dialog.User1Id = dialog.User2Id
		dialog.User2Id = tmp
	}
	id, err := s.repoDialog.CreateDialog(ctx, dialog)

	return id, err
}

func (s *DialogService) GetDialogs(ctx context.Context, userId int) ([]model.Dialog, error) {
	return s.repoDialog.GetDialogs(ctx, userId)
}
