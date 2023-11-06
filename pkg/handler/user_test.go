package handler

import (
	"context"
	"encoding/json"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_user(t *testing.T) {
	mockCookie := &http.Cookie{
		Name:  "session_id",
		Value: "value",
	}
	mockUser := model.User{Mail: "max@max.ru", PasswordHash: "passWord", Name: "Max"}
	response := ClientResponseDto[model.User]{
		Status:  200,
		Message: "success",
		Payload: mockUser,
	}
	jsonData, _ := json.Marshal(response)

	tests := []struct {
		name                 string
		inputCookie          *http.Cookie
		mockBehavior         func(r *mock_service.MockUser)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "Ok",
			inputCookie: mockCookie,
			mockBehavior: func(r *mock_service.MockUser) {
				r.EXPECT().GetCurrentUser(gomock.Any(), 1).Return(mockUser, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: string(jsonData),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockUser(c)
			test.mockBehavior(repo)

			ctx := context.WithValue(context.Background(), keyUserID, 1)
			services := &service.Service{User: repo}
			handler := Handler{services, &ctx}

			mux := http.NewServeMux()
			mux.HandleFunc("/api/v1/user", handler.user)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/api/v1/user", nil)
			if test.inputCookie != nil {
				req.AddCookie(test.inputCookie)
			}

			mux.ServeHTTP(w, req.WithContext(ctx))

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.JSONEq(t, test.expectedResponseBody, w.Body.String())
		})
	}
}
