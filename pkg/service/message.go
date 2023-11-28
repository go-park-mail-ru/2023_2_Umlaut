package service

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
)

type MessageService struct {
	repoMessage repository.Message
}

func NewMessageService(repoMessage repository.Message) *MessageService {
	return &MessageService{repoMessage: repoMessage}
}

func (s *MessageService) GetDialogMessages(ctx context.Context, dialogId int) ([]model.Message, error) {
	return s.repoMessage.GetDialogMessages(ctx, dialogId)
}

func (s *MessageService) SaveOrUpdateMessage(ctx context.Context, message model.Message) (int, error) {
	if message.Id != nil && *message.Id > 0 {
		return s.repoMessage.UpdateMessage(ctx, message)
	}
	return s.repoMessage.CreateMessage(ctx, message)
}
