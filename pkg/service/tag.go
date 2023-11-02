package service

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
)

type TagService struct {
	repoTag repository.Tag
}

func NewTagService(repoTag repository.Tag) *TagService {
	return &TagService{repoTag: repoTag}
}

func (s *TagService) GetAllTags(ctx context.Context) ([]model.Tag, error) {
	return s.repoTag.GetAllTags(ctx)
}

func (s *TagService) GetUserTags(ctx context.Context, userId int) ([]model.Tag, error) {
	return s.repoTag.GetUserTags(ctx, userId)
}
