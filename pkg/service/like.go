package service

import (
	"context"

	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/constants"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/core"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/dto"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
)

type LikeService struct {
	repoLike   repository.Like
	repoDialog repository.Dialog
	repoUser   repository.User
}

func NewLikeService(repoLike repository.Like, repoDialog repository.Dialog, repoUser repository.User) *LikeService {
	return &LikeService{repoLike: repoLike, repoDialog: repoDialog, repoUser: repoUser}
}

func (s *LikeService) CreateLike(ctx context.Context, like core.Like) (core.Dialog, error) {
	_, err := s.repoLike.CreateLike(ctx, like)
	if !like.IsLike || err != nil {
		return core.Dialog{}, err
	}
	mutual, err := s.repoLike.IsMutualLike(ctx, like)
	if err != nil {
		return core.Dialog{}, err
	}
	if mutual {
		dialog := core.Dialog{User1Id: like.LikedByUserId, User2Id: like.LikedToUserId}
		id, err := s.repoDialog.CreateDialog(ctx, dialog)
		if err != nil {
			return core.Dialog{}, err
		}
		dialog.Id = id
		return dialog, constants.ErrMutualLike
	}

	return core.Dialog{}, err
}

func (s *LikeService) GetUserLikedToLikes(ctx context.Context, userId int) (bool, []dto.PremiumLike, error) {
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
