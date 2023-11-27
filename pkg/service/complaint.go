package service

import (
	"context"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository"
)

type ComplaintService struct {
	repoComplaint repository.Complaint
}

func NewComplaintService(repoComplaint repository.Complaint) *ComplaintService {
	return &ComplaintService{repoComplaint: repoComplaint}
}

func (s *ComplaintService) GetComplaintTypes(ctx context.Context) ([]model.ComplaintType, error) {
	return s.repoComplaint.GetComplaintTypes(ctx)
}

func (s *ComplaintService) CreateComplaint(ctx context.Context, complaint model.Complaint) (int, error) {
	return s.repoComplaint.CreateComplaint(ctx, complaint)
}
