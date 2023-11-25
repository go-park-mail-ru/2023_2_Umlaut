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

func (s *AdminService) CreateFeedFeedback(ctx context.Context, rec model.Recommendation) (int, error) {
	return s.repoAdmin.CreateFeedFeedback(ctx, rec)
}

func (s *AdminService) CreateFeedback(ctx context.Context, stat model.Feedback) (int, error) {
	return s.repoAdmin.CreateFeedback(ctx, stat)
}

func (s *AdminService) GetFeedbackStatistics(ctx context.Context) (int, error) {
	feedbacks, err := s.repoAdmin.GetFeedbacks(ctx)
	if err != nil {
		return 0, err
	}
	return 0, nil
}
