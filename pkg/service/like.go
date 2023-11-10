package service

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
)

type LikeService struct {
	repoLike   repository.Like
	repoDialog repository.Dialog
}

func NewLikeService(repoLike repository.Like, repoDialog repository.Dialog) *LikeService {
	return &LikeService{repoLike: repoLike, repoDialog: repoDialog}
}

func (s *LikeService) CreateLike(ctx context.Context, like model.Like) error {
	_, err := s.repoLike.CreateLike(ctx, like)
	if err != nil {
		return err
	}
	mutual, err := s.repoLike.IsMutualLike(ctx, like)
	if err != nil {
		return err
	}
	if mutual {
		dialog := model.Dialog{User1Id: like.LikedByUserId, User2Id: like.LikedToUserId}
		_, err = s.repoDialog.CreateDialog(ctx, dialog)
		if err != nil {
			return err
		}
		return model.MutualLike
	}

	return err
}
