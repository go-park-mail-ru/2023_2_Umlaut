package service

import (
	"context"
	"errors"
	core2 "github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/core"
	"testing"

	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestComplaintService_GetComplaintTypes(t *testing.T) {
	mockComplaintTypes := []core2.ComplaintType{
		{Id: 1, TypeName: "Type1"},
		{Id: 2, TypeName: "Type2"},
	}

	tests := []struct {
		name           string
		mockBehavior   func(r *mock_repository.MockComplaint)
		expectedResult []core2.ComplaintType
		expectedError  error
	}{
		{
			name: "Success",
			mockBehavior: func(r *mock_repository.MockComplaint) {
				r.EXPECT().GetComplaintTypes(gomock.Any()).Return(mockComplaintTypes, nil)
			},
			expectedResult: mockComplaintTypes,
			expectedError:  nil,
		},
		{
			name: "Error",
			mockBehavior: func(r *mock_repository.MockComplaint) {
				r.EXPECT().GetComplaintTypes(gomock.Any()).Return(nil, errors.New("error getting complaint types"))
			},
			expectedResult: nil,
			expectedError:  errors.New("error getting complaint types"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repoComplaint := mock_repository.NewMockComplaint(ctrl)
			test.mockBehavior(repoComplaint)

			service := &ComplaintService{RepoComplaint: repoComplaint}
			complaintTypes, err := service.GetComplaintTypes(context.Background())

			assert.Equal(t, test.expectedResult, complaintTypes)
			assert.Equal(t, test.expectedError, err)
		})
	}
}

func TestComplaintService_CreateComplaint(t *testing.T) {
	mockComplaintID := 1

	tests := []struct {
		name           string
		inputComplaint core2.Complaint
		mockBehavior   func(r *mock_repository.MockComplaint)
		expectedResult int
		expectedError  error
	}{
		{
			name:           "Success",
			inputComplaint: core2.Complaint{ComplaintTypeId: 1},
			mockBehavior: func(r *mock_repository.MockComplaint) {
				r.EXPECT().CreateComplaint(gomock.Any(), gomock.Any()).Return(mockComplaintID, nil)
			},
			expectedResult: mockComplaintID,
			expectedError:  nil,
		},
		{
			name:           "Error",
			inputComplaint: core2.Complaint{},
			mockBehavior: func(r *mock_repository.MockComplaint) {
				r.EXPECT().CreateComplaint(gomock.Any(), gomock.Any()).Return(0, errors.New("error creating complaint"))
			},
			expectedResult: 0,
			expectedError:  errors.New("error creating complaint"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repoComplaint := mock_repository.NewMockComplaint(ctrl)
			test.mockBehavior(repoComplaint)

			service := &ComplaintService{RepoComplaint: repoComplaint}
			complaintID, err := service.CreateComplaint(context.Background(), test.inputComplaint)

			assert.Equal(t, test.expectedResult, complaintID)
			assert.Equal(t, test.expectedError, err)
		})
	}
}

func TestComplaintService_GetNextComplaint(t *testing.T) {
	mockComplaint := core2.Complaint{
		Id:              1,
		ComplaintTypeId: 1,
	}

	tests := []struct {
		name           string
		mockBehavior   func(r *mock_repository.MockComplaint)
		expectedResult core2.Complaint
		expectedError  error
	}{
		{
			name: "Success",
			mockBehavior: func(r *mock_repository.MockComplaint) {
				r.EXPECT().GetNextComplaint(gomock.Any()).Return(mockComplaint, nil)
			},
			expectedResult: mockComplaint,
			expectedError:  nil,
		},
		{
			name: "Error",
			mockBehavior: func(r *mock_repository.MockComplaint) {
				r.EXPECT().GetNextComplaint(gomock.Any()).Return(core2.Complaint{}, errors.New("error getting next complaint"))
			},
			expectedResult: core2.Complaint{},
			expectedError:  errors.New("error getting next complaint"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repoComplaint := mock_repository.NewMockComplaint(ctrl)
			test.mockBehavior(repoComplaint)

			service := &ComplaintService{RepoComplaint: repoComplaint}
			complaint, err := service.GetNextComplaint(context.Background())

			assert.Equal(t, test.expectedResult, complaint)
			assert.Equal(t, test.expectedError, err)
		})
	}
}

func TestComplaintService_AcceptComplaint(t *testing.T) {
	tests := []struct {
		name          string
		complaintID   int
		mockBehavior  func(r *mock_repository.MockComplaint)
		expectedError error
	}{
		{
			name:        "Success",
			complaintID: 1,
			mockBehavior: func(r *mock_repository.MockComplaint) {
				r.EXPECT().AcceptComplaint(gomock.Any(), 1).Return(core2.Complaint{}, nil)
			},
			expectedError: nil,
		},
		{
			name:        "Error",
			complaintID: 0,
			mockBehavior: func(r *mock_repository.MockComplaint) {
				r.EXPECT().AcceptComplaint(gomock.Any(), 0).Return(core2.Complaint{}, errors.New("error accepting complaint"))
			},
			expectedError: errors.New("AcceptComplaint error: error accepting complaint"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repoComplaint := mock_repository.NewMockComplaint(ctrl)
			test.mockBehavior(repoComplaint)

			service := &ComplaintService{RepoComplaint: repoComplaint}
			err := service.AcceptComplaint(context.Background(), test.complaintID)

			assert.Equal(t, test.expectedError, err)
		})
	}
}

func TestComplaintService_DeleteComplaint(t *testing.T) {
	tests := []struct {
		name          string
		complaintID   int
		mockBehavior  func(r *mock_repository.MockComplaint)
		expectedError error
	}{
		{
			name:        "Success",
			complaintID: 1,
			mockBehavior: func(r *mock_repository.MockComplaint) {
				r.EXPECT().DeleteComplaint(gomock.Any(), 1).Return(nil)
			},
			expectedError: nil,
		},
		{
			name:        "Error",
			complaintID: 0,
			mockBehavior: func(r *mock_repository.MockComplaint) {
				r.EXPECT().DeleteComplaint(gomock.Any(), 0).Return(errors.New("error deleting complaint"))
			},
			expectedError: errors.New("error deleting complaint"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repoComplaint := mock_repository.NewMockComplaint(ctrl)
			test.mockBehavior(repoComplaint)

			service := &ComplaintService{RepoComplaint: repoComplaint}
			err := service.DeleteComplaint(context.Background(), test.complaintID)

			assert.Equal(t, test.expectedError, err)
		})
	}
}
