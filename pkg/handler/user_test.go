package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service/mocks"
	"github.com/go-park-mail-ru/2023_2_Umlaut/static"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"mime/multipart"
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
		mockBehavior         func(r *mock_service.MockUser)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Ok",
			mockBehavior: func(r *mock_service.MockUser) {
				r.EXPECT().GetCurrentUser(gomock.Any(), 1).Return(mockUser, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: string(jsonData),
		},
		{
			name: "Error",
			mockBehavior: func(r *mock_service.MockUser) {
				r.EXPECT().GetCurrentUser(gomock.Any(), 1).Return(mockUser, errors.New("some error"))
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"status":500,"message":"some error","payload":""}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockUser(c)
			test.mockBehavior(repo)

			ctx := context.WithValue(context.Background(), static.KeyUserID, 1)
			services := &service.Service{User: repo}
			handler := Handler{services, ctx}

			mux := http.NewServeMux()
			mux.HandleFunc("/api/v1/user", handler.user)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/api/v1/user", nil)
			req.AddCookie(mockCookie)

			mux.ServeHTTP(w, req.WithContext(ctx))

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.JSONEq(t, test.expectedResponseBody, w.Body.String())
		})
	}
}

func TestHandler_updateUser(t *testing.T) {
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
		inputBody            string
		inputUser            model.User
		mockBehavior         func(r *mock_service.MockUser)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"mail": "max@max.ru", "name": "Max", "password": "passWord"}`,
			inputUser: model.User{
				Mail:         "max@max.ru",
				Name:         "Max",
				PasswordHash: "passWord",
			},
			mockBehavior: func(r *mock_service.MockUser) {
				r.EXPECT().UpdateUser(gomock.Any(), mockUser).Return(mockUser, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: string(jsonData),
		},
		{
			name:                 "invalid input body",
			inputBody:            `"mail": "max@max.ru", "name": "Max", "password": "passWord}`,
			inputUser:            model.User{},
			mockBehavior:         func(r *mock_service.MockUser) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"status":400,"message":"invalid input body","payload":""}`,
		},
		{
			name:      "Already Exists",
			inputBody: `{"mail": "max@max.ru", "name": "Max", "password": "passWord"}`,
			inputUser: model.User{
				Mail:         "max@max.ru",
				Name:         "Max",
				PasswordHash: "passWord",
			},
			mockBehavior: func(r *mock_service.MockUser) {
				r.EXPECT().UpdateUser(gomock.Any(), mockUser).Return(mockUser, static.ErrAlreadyExists)
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"status":400,"message":"account with this email already exists","payload":""}`,
		},
		{
			name:      "Invalid User",
			inputBody: `{"mail": "max@max.ru", "name": "Max", "password": "passWord"}`,
			inputUser: model.User{
				Mail:         "max@max.ru",
				Name:         "Max",
				PasswordHash: "passWord",
			},
			mockBehavior: func(r *mock_service.MockUser) {
				r.EXPECT().UpdateUser(gomock.Any(), mockUser).Return(mockUser, static.ErrInvalidUser)
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"status":400,"message":"invalid field for update","payload":""}`,
		},
		{
			name:      "Error",
			inputBody: `{"mail": "max@max.ru", "name": "Max", "password": "passWord"}`,
			inputUser: model.User{
				Mail:         "max@max.ru",
				Name:         "Max",
				PasswordHash: "passWord",
			},
			mockBehavior: func(r *mock_service.MockUser) {
				r.EXPECT().UpdateUser(gomock.Any(), mockUser).Return(mockUser, errors.New("some error"))
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"status":500,"message":"some error","payload":""}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockUser(c)
			test.mockBehavior(repo)

			ctx := context.WithValue(context.Background(), static.KeyUserID, 0)
			services := &service.Service{User: repo}
			handler := Handler{services, ctx}

			mux := http.NewServeMux()
			mux.HandleFunc("/api/v1/user", handler.updateUser)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/user", bytes.NewBufferString(test.inputBody))
			req.AddCookie(mockCookie)

			mux.ServeHTTP(w, req.WithContext(ctx))

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.JSONEq(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_updateUserPhoto(t *testing.T) {
	mockCookie := &http.Cookie{
		Name:  "session_id",
		Value: "value",
	}

	tests := []struct {
		name                 string
		formData             map[string]string
		mockBehavior         func(r *mock_service.MockUser)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Invalid Input Body",
			formData: map[string]string{
				"invalid_key": "testdata/test_image.jpg",
			},
			mockBehavior:         func(r *mock_service.MockUser) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"status":400,"message":"invalid input body","payload":""}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockUser(c)
			test.mockBehavior(repo)

			ctx := context.WithValue(context.Background(), static.KeyUserID, 1)
			services := &service.Service{User: repo}
			handler := Handler{services, ctx}

			mux := http.NewServeMux()
			mux.HandleFunc("/api/v1/user/photo", handler.updateUserPhoto)

			w := httptest.NewRecorder()
			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)
			for key, value := range test.formData {
				_ = writer.WriteField(key, value)
			}
			writer.Close()

			req := httptest.NewRequest("POST", "/api/v1/user/photo", body)
			req.Header.Set("Content-Type", writer.FormDataContentType())
			req.AddCookie(mockCookie)

			mux.ServeHTTP(w, req.WithContext(ctx))

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.JSONEq(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
