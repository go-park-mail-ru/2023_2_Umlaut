package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_createLike(t *testing.T) {
	mockLike := model.Like{
		LikedByUserId: 1,
		LikedToUserId: 2,
	}
	likeJSON, _ := json.Marshal(mockLike)

	tests := []struct {
		name                 string
		requestBody          []byte
		mockBehavior         func(r *mock_service.MockLike, m *mock_service.MockDialog)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "Like Exists",
			requestBody: likeJSON,
			mockBehavior: func(r *mock_service.MockLike, m *mock_service.MockDialog) {
				r.EXPECT().IsLikeExists(gomock.Any(), mockLike).Return(true, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":200,"message":"success","payload":""}`,
		},
		{
			name:        "Error in IsLikeExists",
			requestBody: likeJSON,
			mockBehavior: func(r *mock_service.MockLike, m *mock_service.MockDialog) {
				r.EXPECT().IsLikeExists(gomock.Any(), mockLike).Return(false, errors.New("some error"))
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"status":500,"message":"some error","payload":""}`,
		},
		{
			name:        "Error in CreateLike",
			requestBody: likeJSON,
			mockBehavior: func(r *mock_service.MockLike, m *mock_service.MockDialog) {
				r.EXPECT().IsLikeExists(gomock.Any(), mockLike).Return(false, nil)
				r.EXPECT().CreateLike(gomock.Any(), mockLike).Return(errors.New("some error"))
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"status":500,"message":"some error","payload":""}`,
		},
		{
			name:        "Error in IsUserLiked",
			requestBody: likeJSON,
			mockBehavior: func(r *mock_service.MockLike, m *mock_service.MockDialog) {
				r.EXPECT().IsLikeExists(gomock.Any(), mockLike).Return(false, nil)
				r.EXPECT().CreateLike(gomock.Any(), mockLike).Return(nil)
				r.EXPECT().IsUserLiked(gomock.Any(), mockLike).Return(false, errors.New("some error"))
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"status":500,"message":"some error","payload":""}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repoLike := mock_service.NewMockLike(c)
			repoDialog := mock_service.NewMockDialog(c)
			test.mockBehavior(repoLike, repoDialog)

			ctx := context.WithValue(context.Background(), keyUserID, 1)
			services := &service.Service{Like: repoLike, Dialog: repoDialog}
			handler := Handler{services, &ctx}

			mux := http.NewServeMux()
			mux.HandleFunc("/api/v1/like", handler.createLike)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/like", bytes.NewReader(test.requestBody))
			mux.ServeHTTP(w, req.WithContext(ctx))

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
