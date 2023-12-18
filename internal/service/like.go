package service

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/constants"
	core2 "github.com/go-park-mail-ru/2023_2_Umlaut/internal/model/core"
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/repository"
)

type LikeService struct {
	repoLike   repository.Like
	repoDialog repository.Dialog
	repoUser   repository.User
}

func NewLikeService(repoLike repository.Like, repoDialog repository.Dialog, repoUser repository.User) *LikeService {
	return &LikeService{repoLike: repoLike, repoDialog: repoDialog, repoUser: repoUser}
}

func (s *LikeService) CreateLike(ctx context.Context, like core2.Like) (core2.Dialog, error) {
	_, err := s.repoLike.CreateLike(ctx, like)
	if err != nil {
		return core2.Dialog{}, err
	}
	if !like.IsLike {
		return core2.Dialog{}, nil
	}
	mutual, err := s.repoLike.IsMutualLike(ctx, like)
	if err != nil {
		return core2.Dialog{}, err
	}
	if mutual {
		id, err := s.repoDialog.CreateDialog(ctx, core2.Dialog{User1Id: like.LikedByUserId, User2Id: like.LikedToUserId})
		if err != nil {
			return core2.Dialog{}, err
		}
		dialog, err := s.repoDialog.GetDialogById(ctx, id)
		if err != nil {
			return core2.Dialog{}, err
		}
		return dialog, constants.ErrMutualLike
	}

	return core2.Dialog{}, err
}

func (s *LikeService) GetUserLikedToLikes(ctx context.Context, userId int) (bool, []core2.PremiumLike, error) {
	user, err := s.repoUser.GetUserById(ctx, userId)
	if err != nil {
		return false, nil, err
	}
	likes, err := s.repoLike.GetUserLikedToLikes(ctx, userId)
	if err != nil {
		return false, nil, err
	}
	if user.Role != 2 {
		return false, likes, nil
	}
	return true, likes, nil
}
