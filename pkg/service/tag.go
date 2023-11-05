package service

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
)

type TagService struct {
	repoTag repository.Tag
}

func NewTagService(repoTag repository.Tag) *TagService {
	return &TagService{repoTag: repoTag}
}

func (s *TagService) GetAllTags(ctx context.Context) ([]string, error) {
	return s.repoTag.GetAllTags(ctx)
}
