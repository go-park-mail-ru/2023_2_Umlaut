package service

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/model/core"
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/repository"
)

type DialogService struct {
	repoDialog repository.Dialog
}

func NewDialogService(repoDialog repository.Dialog) *DialogService {
	return &DialogService{repoDialog: repoDialog}
}

func (s *DialogService) CreateDialog(ctx context.Context, dialog core.Dialog) (int, error) {
	return s.repoDialog.CreateDialog(ctx, dialog)
}

func (s *DialogService) GetDialogs(ctx context.Context, userId int) ([]core.Dialog, error) {
	return s.repoDialog.GetDialogs(ctx, userId)
}

func (s *DialogService) GetDialog(ctx context.Context, id int) (core.Dialog, error) {
	return s.repoDialog.GetDialogById(ctx, id)
}
