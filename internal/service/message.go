package service

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/model/core"
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/repository"
)

type MessageService struct {
	repoMessage repository.Message
}

func NewMessageService(repoMessage repository.Message) *MessageService {
	return &MessageService{repoMessage: repoMessage}
}

func (s *MessageService) GetDialogMessages(ctx context.Context, userId int, recipientId int) ([]core.Message, error) {
	return s.repoMessage.GetDialogMessages(ctx, userId, recipientId)
}

func (s *MessageService) SaveOrUpdateMessage(ctx context.Context, message core.Message) (core.Message, error) {
	if message.Id != nil && *message.Id > 0 {
		return s.repoMessage.UpdateMessage(ctx, message)
	}
	return s.repoMessage.CreateMessage(ctx, message)
}
