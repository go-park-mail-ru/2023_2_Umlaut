package service

import (
	"context"
	"fmt"
	core2 "github.com/go-park-mail-ru/2023_2_Umlaut/internal/model/core"
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/repository"
)

type ComplaintService struct {
	repoComplaint repository.Complaint
}

func NewComplaintService(repoComplaint repository.Complaint) *ComplaintService {
	return &ComplaintService{repoComplaint: repoComplaint}
}

func (s *ComplaintService) GetComplaintTypes(ctx context.Context) ([]core2.ComplaintType, error) {
	return s.repoComplaint.GetComplaintTypes(ctx)
}

func (s *ComplaintService) CreateComplaint(ctx context.Context, complaint core2.Complaint) (int, error) {
	return s.repoComplaint.CreateComplaint(ctx, complaint)
}

func (s *ComplaintService) GetNextComplaint(ctx context.Context) (core2.Complaint, error) {
	return s.repoComplaint.GetNextComplaint(ctx)
}

func (s *ComplaintService) AcceptComplaint(ctx context.Context, complaintId int) error {
	_, err := s.repoComplaint.AcceptComplaint(ctx, complaintId)
	if err != nil {
		return fmt.Errorf("AcceptComplaint error: %v", err)
	}

	return nil
}

func (s *ComplaintService) DeleteComplaint(ctx context.Context, complaintId int) error {
	return s.repoComplaint.DeleteComplaint(ctx, complaintId)
}
