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
	mutual, err := s.repoLike.IsMutualLike(ctx, like)
	if err != nil {
		return err
	}
	if mutual {
		dialog := model.Dialog{User1Id: like.LikedByUserId, User2Id: like.LikedToUserId}
		_, errDialog := s.repoDialog.CreateDialog(ctx, dialog)
		if errDialog != nil {
			return errDialog
		}
		return model.MutualLike
	}

	_, err = s.repoLike.CreateLike(ctx, like)
	return err
}
