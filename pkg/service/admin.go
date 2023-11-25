package service

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
)

type AdminService struct {
	repoAdmin repository.Admin
}

func NewAdminService(repoAdmin repository.Admin) *AdminService {
	return &AdminService{repoAdmin: repoAdmin}
}

func (s *AdminService) GetStatistic(ctx context.Context) int {
	return 0
}
