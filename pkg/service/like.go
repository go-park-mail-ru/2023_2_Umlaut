package service

import (
	"context"
	"errors"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
)

type LikeService struct {
	repoLike repository.Like
}

func NewLikeService(repoLike repository.Like) *LikeService {
	return &LikeService{repoLike: repoLike}
}

func (s *LikeService) CreateLike(ctx context.Context, like model.Like) error {
	exists, err := s.repoLike.Exists(ctx, like)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("this like is already exists")
	}
	
	_, err = s.repoLike.CreateLike(ctx, like)

	return err
}

func (s *LikeService) IsUserLiked(ctx context.Context, like model.Like) (bool, error) {
	tmp := like.LikedByUserId
	like.LikedByUserId = like.LikedToUserId
	like.LikedToUserId = tmp

	exists, err := s.repoLike.Exists(ctx, like)

	return exists, err
}
