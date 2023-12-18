package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	static2 "github.com/go-park-mail-ru/2023_2_Umlaut/pkg/constants"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/core"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/dto"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_user(t *testing.T) {
	mockCookie := &http.Cookie{
		Name:  "session_id",
		Value: "value",
	}
	mockUser := core.User{Mail: "user@user.ru", PasswordHash: "passWord", Name: "user"}
	response := dto.ClientResponseDto{
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
			name: "Banned user",
			mockBehavior: func(r *mock_service.MockUser) {
				r.EXPECT().GetCurrentUser(gomock.Any(), 1).Return(mockUser, static2.ErrBannedUser)
			},
			expectedStatusCode:   http.StatusForbidden,
			expectedResponseBody: `{"status":403,"message":"this user is blocked","payload":""}`,
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

			ctx := context.WithValue(context.Background(), static2.KeyUserID, 1)
			services := &service.Service{User: repo}
			handler := Handler{services: services}

			mux := http.NewServeMux()
			mux.HandleFunc("/api/v1/user", handler.user)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/api/v1/user", nil)
			req.AddCookie(mockCookie)

			mux.ServeHTTP(w, req.WithContext(ctx))

			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_updateUser(t *testing.T) {
	mockCookie := &http.Cookie{
		Name:  "session_id",
		Value: "value",
	}
	mockUser := core.User{Id: 1, Mail: "user@user.ru", PasswordHash: "passWord", Name: "user"}
	response := dto.ClientResponseDto{
		Status:  200,
		Message: "success",
		Payload: mockUser,
	}
	jsonData, _ := json.Marshal(response)

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            core.User
		mockBehavior         func(r *mock_service.MockUser)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"mail": "user@user.ru", "name": "user", "password": "passWord"}`,
			inputUser: core.User{
				Mail:         "user@user.ru",
				Name:         "user",
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
			inputBody:            `"mail": "user@user.ru", "name": "user", "password": "passWord}`,
			inputUser:            core.User{},
			mockBehavior:         func(r *mock_service.MockUser) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"status":400,"message":"invalid input body","payload":""}`,
		},
		{
			name:      "Already Exists",
			inputBody: `{"mail": "user@user.ru", "name": "user", "password": "passWord"}`,
			inputUser: core.User{
				Mail:         "user@user.ru",
				Name:         "user",
				PasswordHash: "passWord",
			},
			mockBehavior: func(r *mock_service.MockUser) {
				r.EXPECT().UpdateUser(gomock.Any(), mockUser).Return(mockUser, static2.ErrAlreadyExists)
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"status":400,"message":"account with this email already exists","payload":""}`,
		},
		{
			name:      "Invalid User",
			inputBody: `{"mail": "user@user.ru", "name": "user", "password": "passWord"}`,
			inputUser: core.User{
				Mail:         "user@user.ru",
				Name:         "user",
				PasswordHash: "passWord",
			},
			mockBehavior: func(r *mock_service.MockUser) {
				r.EXPECT().UpdateUser(gomock.Any(), mockUser).Return(mockUser, static2.ErrInvalidUser)
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"status":400,"message":"invalid field for update","payload":""}`,
		},
		{
			name:      "Error",
			inputBody: `{"mail": "user@user.ru", "name": "user", "password": "passWord"}`,
			inputUser: core.User{
				Mail:         "user@user.ru",
				Name:         "user",
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

			ctx := context.WithValue(context.Background(), static2.KeyUserID, 1)
			services := &service.Service{User: repo}
			handler := Handler{services: services}

			mux := http.NewServeMux()
			mux.HandleFunc("/api/v1/user", handler.updateUser)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/user", bytes.NewBufferString(test.inputBody))
			req.AddCookie(mockCookie)

			mux.ServeHTTP(w, req.WithContext(ctx))
			r := w.Body.String()
			assert.Equal(t, r, test.expectedResponseBody)
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

			ctx := context.WithValue(context.Background(), static2.KeyUserID, 1)
			services := &service.Service{User: repo}
			handler := Handler{services: services}

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

			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_deleteUserPhoto(t *testing.T) {
	mockUserID := 1
	mockLink := "photo.jpg"
	mockCookie := &http.Cookie{
		Name:  "session_id",
		Value: "value",
	}
	response := dto.ClientResponseDto{
		Status:  200,
		Message: "success",
		Payload: "",
	}
	jsonData, _ := json.Marshal(response)

	tests := []struct {
		name                 string
		inputBody            string
		mockBehavior         func(r *mock_service.MockUser)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"link": "photo.jpg"}`,
			mockBehavior: func(r *mock_service.MockUser) {
				r.EXPECT().DeleteFile(gomock.Any(), mockUserID, mockLink).Return(nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: string(jsonData),
		},
		{
			name:      "banned user",
			inputBody: `{"link": "photo.jpg"}`,
			mockBehavior: func(r *mock_service.MockUser) {
				r.EXPECT().DeleteFile(gomock.Any(), mockUserID, mockLink).Return(static2.ErrBannedUser)
			},
			expectedStatusCode:   http.StatusForbidden,
			expectedResponseBody: `{"status":403,"message":"this user is blocked","payload":""}`,
		},
		{
			name:      "No photo",
			inputBody: `{"link": "photo.jpg"}`,
			mockBehavior: func(r *mock_service.MockUser) {
				r.EXPECT().DeleteFile(gomock.Any(), mockUserID, mockLink).Return(static2.ErrNoFiles)
			},
			expectedStatusCode:   http.StatusNotFound,
			expectedResponseBody: `{"status":404,"message":"This user has no photos","payload":""}`,
		},
		{
			name:      "Error",
			inputBody: `{"link": "photo.jpg"}`,
			mockBehavior: func(r *mock_service.MockUser) {
				r.EXPECT().DeleteFile(gomock.Any(), mockUserID, mockLink).Return(errors.New("some error"))
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

			ctx := context.WithValue(context.Background(), static2.KeyUserID, 1)
			services := &service.Service{User: repo}
			handler := Handler{services: services}

			mux := http.NewServeMux()
			mux.HandleFunc("/api/v1/user/photo", handler.deleteUserPhoto)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("DELETE", "/api/v1/user/photo", bytes.NewBufferString(test.inputBody))
			req.AddCookie(mockCookie)

			mux.ServeHTTP(w, req.WithContext(ctx))

			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
