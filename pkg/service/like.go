package service

import (
	"context"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
	"github.com/go-park-mail-ru/2023_2_Umlaut/static"
)

type LikeService struct {
	repoLike   repository.Like
	repoDialog repository.Dialog
	repoUser   repository.User
}

func NewLikeService(repoLike repository.Like, repoDialog repository.Dialog, repoUser repository.User) *LikeService {
	return &LikeService{repoLike: repoLike, repoDialog: repoDialog, repoUser: repoUser}
}

func (s *LikeService) CreateLike(ctx context.Context, like model.Like) (model.Dialog, error) {
	_, err := s.repoLike.CreateLike(ctx, like)
	if err != nil {
		return model.Dialog{}, err
	}
	if !like.IsLike {
		return model.Dialog{}, nil
	}
	mutual, err := s.repoLike.IsMutualLike(ctx, like)
	if err != nil {
		return model.Dialog{}, err
	}
	if mutual {
		id, err := s.repoDialog.CreateDialog(ctx, model.Dialog{User1Id: like.LikedByUserId, User2Id: like.LikedToUserId})
		if err != nil {
			return model.Dialog{}, err
		}
		dialog, err := s.repoDialog.GetDialogById(ctx, id)
		if err != nil {
			return model.Dialog{}, err
		}
		return dialog, static.ErrMutualLike
	}

	return model.Dialog{}, err
}

func (s *LikeService) GetUserLikedToLikes(ctx context.Context, userId int) ([]model.PremiumLike, error) {
	user, err := s.repoUser.GetUserById(ctx, userId)
	if err != nil {
		return nil, err
	}
	if user.Role != 2 {
		return nil, static.ErrNoAccess
	}

	return s.repoLike.GetUserLikedToLikes(ctx, userId)
}
