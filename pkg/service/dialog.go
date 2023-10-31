package service

import (
	"context"
	"fmt"

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
	exists, err := s.repoDialog.Exists(ctx, dialog)
	if err != nil {
		return 0, err
	}
	if exists {
		return 0, fmt.Errorf("dialog is already exists")
	}

	id, err := s.repoDialog.CreateDialog(ctx, dialog)

	return id, err
}

func (s *DialogService) GetDialogs(ctx context.Context, userId int) ([]model.Dialog, error) {
	dialogs, err := s.repoDialog.GetDialogs(ctx, userId)

	return dialogs, err
}
