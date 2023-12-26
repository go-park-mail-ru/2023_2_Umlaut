package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/constants"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/core"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_createLike(t *testing.T) {
	mockLike := core.Like{
		LikedByUserId: 1,
		LikedToUserId: 2,
	}
	likeJSON, _ := json.Marshal(mockLike)
	tests := []struct {
		name                 string
		requestBody          string
		mockBehavior         func(r *mock_service.MockLike, m *mock_service.MockDialog)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "Like Create",
			requestBody: string(likeJSON),
			mockBehavior: func(r *mock_service.MockLike, m *mock_service.MockDialog) {
				r.EXPECT().CreateLike(gomock.Any(), mockLike).Return(core.Dialog{}, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":200,"message":"success","payload":""}`,
		},
		{
			name:        "already liked",
			requestBody: string(likeJSON),
			mockBehavior: func(r *mock_service.MockLike, m *mock_service.MockDialog) {
				r.EXPECT().CreateLike(gomock.Any(), mockLike).Return(core.Dialog{}, constants.ErrAlreadyExists)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":200,"message":"already liked","payload":""}`,
		},
		{
			name:                 "invalid input body",
			requestBody:          `{"liked_by_user_id","liked_to_user_id":2}`,
			mockBehavior:         func(r *mock_service.MockLike, m *mock_service.MockDialog) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"status":400,"message":"invalid input body","payload":""}`,
		},
		{
			name:        "Error",
			requestBody: string(likeJSON),
			mockBehavior: func(r *mock_service.MockLike, m *mock_service.MockDialog) {
				r.EXPECT().CreateLike(gomock.Any(), mockLike).Return(core.Dialog{}, errors.New("some error"))
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

			ctx := context.WithValue(context.Background(), constants.KeyUserID, 1)
			services := &service.Service{Like: repoLike, Dialog: repoDialog}
			handler := Handler{services: services}

			mux := http.NewServeMux()
			mux.HandleFunc("/api/v1/like", handler.createLike)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/like", bytes.NewBufferString(test.requestBody))
			mux.ServeHTTP(w, req.WithContext(ctx))

			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
