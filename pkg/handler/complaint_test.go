package handler

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service"
	mock_service "github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service/mocks"
	"github.com/go-park-mail-ru/2023_2_Umlaut/static"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_createComplaint(t *testing.T) {

	mockComplaint := model.Complaint{
		ReporterUserId:  1,
		ReportedUserId:  2,
		ComplaintTypeId: 1,
	}

	tests := []struct {
		name                 string
		requestBody          string
		mockBehavior         func(r *mock_service.MockComplaint)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "Ok",
			requestBody: `{"reported_user_id": 2, "complaint_type_id": 1}`,
			mockBehavior: func(r *mock_service.MockComplaint) {
				r.EXPECT().CreateComplaint(gomock.Any(), mockComplaint).Return(1, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":200,"message":"success","payload":""}`,
		},
		{
			name:        "invalid input body",
			requestBody: `{"reported_user_id": 2, "complaint_type_id": "complaint}`,
			mockBehavior: func(r *mock_service.MockComplaint) {
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"status":400,"message":"invalid input body","payload":""}`,
		},
		{
			name:        "already exists",
			requestBody: `{"reported_user_id": 2, "complaint_type_id": 1}`,
			mockBehavior: func(r *mock_service.MockComplaint) {
				r.EXPECT().CreateComplaint(gomock.Any(), mockComplaint).Return(1, static.ErrAlreadyExists)
			},
			expectedStatusCode:   http.StatusConflict,
			expectedResponseBody: `{"status":409,"message":"complaint already exists","payload":""}`,
		},
		{
			name:        "error",
			requestBody: `{"reported_user_id": 2, "complaint_type_id": 1}`,
			mockBehavior: func(r *mock_service.MockComplaint) {
				r.EXPECT().CreateComplaint(gomock.Any(), mockComplaint).Return(1, errors.New("some error"))
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"status":500,"message":"some error","payload":""}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repoComplaint := mock_service.NewMockComplaint(c)
			test.mockBehavior(repoComplaint)

			ctx := context.WithValue(context.Background(), static.KeyUserID, 1)
			services := &service.Service{Complaint: repoComplaint}
			handler := Handler{services: services}

			mux := http.NewServeMux()
			mux.HandleFunc("/api/v1/complaint", handler.createComplaint)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/complaint", bytes.NewBufferString(test.requestBody))
			mux.ServeHTTP(w, req.WithContext(ctx))

			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_getAllComplaintTypes(t *testing.T) {
	tests := []struct {
		name                 string
		mockBehavior         func(r *mock_service.MockComplaint)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Ok",
			mockBehavior: func(r *mock_service.MockComplaint) {
				complaintTypes := []model.ComplaintType{
					{
						Id:       1,
						TypeName: "type1",
					},
					{
						Id:       2,
						TypeName: "type2",
					},
				}
				r.EXPECT().GetComplaintTypes(gomock.Any()).Return(complaintTypes, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":200,"message":"success","payload":[{"id":1,"type_name":"type1"},{"id":2,"type_name":"type2"}]}`,
		},
		{
			name: "error",
			mockBehavior: func(r *mock_service.MockComplaint) {
				r.EXPECT().GetComplaintTypes(gomock.Any()).Return(nil, errors.New("some error"))
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"status":500,"message":"some error","payload":""}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			complaintService := mock_service.NewMockComplaint(c)
			test.mockBehavior(complaintService)

			ctx := context.WithValue(context.Background(), static.KeyUserID, 1)
			services := &service.Service{Complaint: complaintService}
			handler := Handler{services: services}

			mux := http.NewServeMux()
			mux.HandleFunc("/api/v1/complaint_types", handler.getAllComplaintTypes)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/api/v1/complaint_types", nil)

			mux.ServeHTTP(w, req.WithContext(ctx))

			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
