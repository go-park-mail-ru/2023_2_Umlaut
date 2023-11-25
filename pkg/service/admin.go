package service

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
)

type AdminService struct {
	repoAdmin repository.Admin
}

func NewAdminService(repoAdmin repository.Admin) *AdminService {
	return &AdminService{repoAdmin: repoAdmin}
}

func (s *AdminService) GetStatistic(ctx context.Context) (int, error) {
	return 0, nil
}

func (s *AdminService) CreateRecommendation(ctx context.Context, rec model.Recommendation) (int, error) {
	return s.repoAdmin.CreateRecommendation(ctx, rec)
}

func (s *AdminService) CreateStatistic(ctx context.Context, stat model.Statistic) (int, error) {
	return 0, nil
}
