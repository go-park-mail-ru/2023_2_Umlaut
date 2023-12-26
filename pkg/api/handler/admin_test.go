package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/constants"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/admin/proto"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/admin/proto/mocks"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/core"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service/mocks"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/utils"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"net/http/httptest"
	"testing"
)

import (
	"context"
	"github.com/stretchr/testify/assert"
)

func TestHandler_createFeedback(t *testing.T) {
	mockFeedback := &proto.Feedback{
		UserId:  1,
		Rating:  2,
		Liked:   "yes",
		NeedFix: "no",
		Comment: "hello",
	}
	mockFeedbackInput := core.Feedback{
		UserId:  1,
		Rating:  utils.ToPtrInt(2),
		Liked:   utils.ToPtrString("yes"),
		NeedFix: utils.ToPtrString("no"),
		Comment: utils.ToPtrString("hello"),
	}
	ctx := context.WithValue(context.Background(), constants.KeyUserID, 1)

	feedbackJSON, _ := json.Marshal(mockFeedbackInput)

	tests := []struct {
		name                 string
		requestBody          string
		mockBehavior         func(r *mock_proto.MockAdminClient)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "Feedback Create",
			requestBody: string(feedbackJSON),
			mockBehavior: func(r *mock_proto.MockAdminClient) {
				r.EXPECT().CreateFeedback(ctx, mockFeedback).Return(nil, nil)
			},
			expectedResponseBody: `{"status":200,"message":"success","payload":""}`,
		},
		{
			name:        "Feedback Create Error InvalidArgument",
			requestBody: string(feedbackJSON),
			mockBehavior: func(r *mock_proto.MockAdminClient) {
				r.EXPECT().CreateFeedback(ctx, mockFeedback).Return(
					nil,
					status.Errorf(codes.InvalidArgument, "Неверный аргумент"),
				)
			},
			expectedResponseBody: `{"status":400,"message":"Неверный аргумент","payload":""}`,
		},
		{
			name:        "Feedback Create InternalServerError",
			requestBody: string(feedbackJSON),
			mockBehavior: func(r *mock_proto.MockAdminClient) {
				r.EXPECT().CreateFeedback(ctx, mockFeedback).Return(
					nil,
					errors.New("some error"),
				)
			},
			expectedResponseBody: `{"status":500,"message":"some error","payload":""}`,
		},
		{
			name:                 "Invalid input body",
			requestBody:          `{"Rating":4}`,
			mockBehavior:         func(r *mock_proto.MockAdminClient) {},
			expectedResponseBody: `{"status":400,"message":"invalid input body","payload":""}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			adminClient := mock_proto.NewMockAdminClient(c)
			test.mockBehavior(adminClient)

			handler := Handler{adminMicroservice: adminClient}

			mux := http.NewServeMux()
			mux.HandleFunc("/api/v1/feedback", handler.createFeedback)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/feedback", bytes.NewBufferString(test.requestBody))
			mux.ServeHTTP(w, req.WithContext(ctx))

			assert.Equal(t, test.expectedResponseBody, w.Body.String())
		})
	}
}

func TestHandler_createRecommendation(t *testing.T) {
	mockRecommendation := &proto.Recommendation{
		UserId: 1,
		Rating: 2,
	}
	mockRecommendationInput := core.Recommendation{
		UserId: 1,
		Rating: utils.ToPtrInt(2),
	}
	ctx := context.WithValue(context.Background(), constants.KeyUserID, 1)

	recommendationJSON, _ := json.Marshal(mockRecommendationInput)

	tests := []struct {
		name                 string
		requestBody          string
		mockBehavior         func(r *mock_proto.MockAdminClient)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "Recommendation Create",
			requestBody: string(recommendationJSON),
			mockBehavior: func(r *mock_proto.MockAdminClient) {
				r.EXPECT().CreateRecommendation(ctx, mockRecommendation).Return(nil, nil)
			},
			expectedResponseBody: `{"status":200,"message":"success","payload":""}`,
		},
		{
			name:        "Recommendation Create Error InvalidArgument",
			requestBody: string(recommendationJSON),
			mockBehavior: func(r *mock_proto.MockAdminClient) {
				r.EXPECT().CreateRecommendation(ctx, mockRecommendation).Return(
					nil,
					status.Errorf(codes.Internal, "упало"),
				)
			},
			expectedResponseBody: `{"status":500,"message":"упало","payload":""}`,
		},
		{
			name:        "Recommendation Create InternalServerError",
			requestBody: string(recommendationJSON),
			mockBehavior: func(r *mock_proto.MockAdminClient) {
				r.EXPECT().CreateRecommendation(ctx, mockRecommendation).Return(
					nil,
					errors.New("some error"),
				)
			},
			expectedResponseBody: `{"status":500,"message":"some error","payload":""}`,
		},
		{
			name:                 "Invalid input body",
			requestBody:          `{}`,
			mockBehavior:         func(r *mock_proto.MockAdminClient) {},
			expectedResponseBody: `{"status":400,"message":"invalid input body","payload":""}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			adminClient := mock_proto.NewMockAdminClient(c)
			test.mockBehavior(adminClient)

			handler := Handler{adminMicroservice: adminClient}

			mux := http.NewServeMux()
			mux.HandleFunc("/api/v1/recommendation", handler.createRecommendation)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/recommendation", bytes.NewBufferString(test.requestBody))
			mux.ServeHTTP(w, req.WithContext(ctx))

			assert.Equal(t, test.expectedResponseBody, w.Body.String())
		})
	}
}

func TestHandler_createFeedFeedback(t *testing.T) {
	mockRecommendation := &proto.Recommendation{
		UserId: 1,
		Rating: 2,
	}
	mockRecommendationInput := core.Recommendation{
		UserId: 1,
		Rating: utils.ToPtrInt(2),
	}
	ctx := context.WithValue(context.Background(), constants.KeyUserID, 1)

	recommendationJSON, _ := json.Marshal(mockRecommendationInput)

	tests := []struct {
		name                 string
		requestBody          string
		mockBehavior         func(r *mock_proto.MockAdminClient)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "FeedFeedback Create",
			requestBody: string(recommendationJSON),
			mockBehavior: func(r *mock_proto.MockAdminClient) {
				r.EXPECT().CreateRecommendation(ctx, mockRecommendation).Return(nil, nil)
			},
			expectedResponseBody: `{"status":200,"message":"success","payload":""}`,
		},
		{
			name:        "FeedFeedback Create Error InvalidArgument",
			requestBody: string(recommendationJSON),
			mockBehavior: func(r *mock_proto.MockAdminClient) {
				r.EXPECT().CreateRecommendation(ctx, mockRecommendation).Return(
					nil,
					status.Errorf(codes.Unauthenticated, "нельзя"),
				)
			},
			expectedResponseBody: `{"status":401,"message":"нельзя","payload":""}`,
		},
		{
			name:        "FeedFeedback Create InternalServerError",
			requestBody: string(recommendationJSON),
			mockBehavior: func(r *mock_proto.MockAdminClient) {
				r.EXPECT().CreateRecommendation(ctx, mockRecommendation).Return(
					nil,
					errors.New("some error"),
				)
			},
			expectedResponseBody: `{"status":500,"message":"some error","payload":""}`,
		},
		{
			name:                 "Invalid input body",
			requestBody:          `{}`,
			mockBehavior:         func(r *mock_proto.MockAdminClient) {},
			expectedResponseBody: `{"status":400,"message":"invalid input body","payload":""}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			adminClient := mock_proto.NewMockAdminClient(c)
			test.mockBehavior(adminClient)

			handler := Handler{adminMicroservice: adminClient}

			mux := http.NewServeMux()
			mux.HandleFunc("/api/v1/feed-feedback", handler.createFeedFeedback)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/feed-feedback", bytes.NewBufferString(test.requestBody))
			mux.ServeHTTP(w, req.WithContext(ctx))

			assert.Equal(t, test.expectedResponseBody, w.Body.String())
		})
	}
}

func TestHandler_showCSAT(t *testing.T) {
	mockCSATType := 2
	mockID := 1
	ctx := context.WithValue(context.Background(), constants.KeyUserID, mockID)

	tests := []struct {
		name                 string
		mockBehavior         func(r *mock_service.MockAdmin)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Get CSAT Type Success",
			mockBehavior: func(r *mock_service.MockAdmin) {
				r.EXPECT().GetCSATType(gomock.Any(), mockID).Return(mockCSATType, nil)
			},
			expectedResponseBody: `{"status":200,"message":"success","payload":2}`,
		},
		{
			name: "Get CSAT Type Error",
			mockBehavior: func(r *mock_service.MockAdmin) {
				r.EXPECT().GetCSATType(gomock.Any(), mockID).Return(0, errors.New("CSAT type not found"))
			},
			expectedResponseBody: `{"status":500,"message":"CSAT type not found","payload":""}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			adminService := mock_service.NewMockAdmin(c)
			test.mockBehavior(adminService)

			handler := Handler{services: &service.Service{Admin: adminService}}

			mux := http.NewServeMux()
			mux.HandleFunc("/api/v1/show-csat", handler.showCSAT)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/api/v1/show-csat", nil)
			mux.ServeHTTP(w, req.WithContext(ctx))

			assert.Equal(t, test.expectedResponseBody, w.Body.String())
		})
	}
}

func TestHandler_getRecommendationStatistic(t *testing.T) {
	mockRecommendationStatistic := &proto.RecommendationStatistic{
		AvgRecommend:   4.5,
		NPS:            70.0,
		RecommendCount: []int32{10, 20, 30},
	}
	ctx := context.Background()

	tests := []struct {
		name                 string
		mockBehavior         func(r *mock_proto.MockAdminClient)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Get Recommendation Statistic Success",
			mockBehavior: func(r *mock_proto.MockAdminClient) {
				r.EXPECT().GetRecommendationStatistic(ctx, gomock.Any()).Return(mockRecommendationStatistic, nil)
			},
			expectedResponseBody: `{"status":200,"message":"success","payload":{"avg-recommend":4.5,"nps":70,"recommend-count":[10,20,30]}}`,
		},
		{
			name: "Get Recommendation Statistic Error",
			mockBehavior: func(r *mock_proto.MockAdminClient) {
				r.EXPECT().GetRecommendationStatistic(ctx, gomock.Any()).Return(nil, errors.New("failed to fetch recommendation statistic"))
			},
			expectedResponseBody: `{"status":500,"message":"failed to fetch recommendation statistic","payload":""}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			adminClient := mock_proto.NewMockAdminClient(c)
			test.mockBehavior(adminClient)

			handler := Handler{adminMicroservice: adminClient}

			mux := http.NewServeMux()
			mux.HandleFunc("/api/v1/recommendation-statistic", handler.getRecommendationStatistic)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/api/v1/recommendation-statistic", nil)
			mux.ServeHTTP(w, req.WithContext(ctx))

			assert.Equal(t, test.expectedResponseBody, w.Body.String())
		})
	}
}

func TestHandler_getFeedbackStatistic(t *testing.T) {
	mockFeedbackStatistic := &proto.FeedbackStatistic{
		AvgRating:   4.0,
		RatingCount: []int32{5, 10, 15},
		Comments:    []string{"comment1", "comment2"},
	}
	ctx := context.Background()

	tests := []struct {
		name                 string
		mockBehavior         func(r *mock_proto.MockAdminClient)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Get Feedback Statistic Success",
			mockBehavior: func(r *mock_proto.MockAdminClient) {
				r.EXPECT().GetFeedbackStatistic(ctx, gomock.Any()).Return(mockFeedbackStatistic, nil)
			},
			expectedResponseBody: `{"status":200,"message":"success","payload":{"avg-rating":4,"rating-count":[5,10,15],"liked-map":{},"need-fix-map":{},"comments":["comment1","comment2"]}}`,
		},
		{
			name: "Get Feedback Statistic Error",
			mockBehavior: func(r *mock_proto.MockAdminClient) {
				r.EXPECT().GetFeedbackStatistic(ctx, gomock.Any()).Return(nil, errors.New("failed to fetch feedback statistic"))
			},
			expectedResponseBody: `{"status":500,"message":"failed to fetch feedback statistic","payload":""}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			adminClient := mock_proto.NewMockAdminClient(c)
			test.mockBehavior(adminClient)

			handler := Handler{adminMicroservice: adminClient}

			mux := http.NewServeMux()
			mux.HandleFunc("/api/v1/feedback-statistic", handler.getFeedbackStatistic)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/api/v1/feedback-statistic", nil)
			mux.ServeHTTP(w, req.WithContext(ctx))

			assert.Equal(t, test.expectedResponseBody, w.Body.String())
		})
	}
}
