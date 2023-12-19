package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/auth/proto"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/auth/proto/mocks"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/dto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_signUp(t *testing.T) {
	mockSignUpInput := &proto.SignUpInput{
		Mail:      "test@example.com",
		Password:  "password",
		Name:      "Test User",
		InvitedBy: "",
	}
	mockSignUpInputDto := &dto.SignUpInput{
		Mail:      "test@example.com",
		Password:  "password",
		Name:      "Test User",
		InvitedBy: nil,
	}
	mockIdResponse := &proto.UserId{Id: 1, Cookie: &proto.Cookie{}}

	signUpInputJSON, _ := json.Marshal(mockSignUpInputDto)
	ctx := context.Background()

	tests := []struct {
		name                 string
		requestBody          string
		mockBehavior         func(r *mock_proto.MockAuthorizationClient)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "Sign Up Success",
			requestBody: string(signUpInputJSON),
			mockBehavior: func(r *mock_proto.MockAuthorizationClient) {
				r.EXPECT().SignUp(ctx, mockSignUpInput).Return(mockIdResponse, nil)
			},
			expectedResponseBody: `{"status":200,"message":"success","payload":{"id":1}}`,
		},
		{
			name:        "Sign Up Error InvalidArgument",
			requestBody: string(signUpInputJSON),
			mockBehavior: func(r *mock_proto.MockAuthorizationClient) {
				r.EXPECT().SignUp(ctx, mockSignUpInput).Return(
					nil,
					status.Errorf(codes.InvalidArgument, "Неверный аргумент"),
				)
			},
			expectedResponseBody: `{"status":400,"message":"Неверный аргумент","payload":""}`,
		},
		{
			name:        "Sign Up InternalServerError",
			requestBody: string(signUpInputJSON),
			mockBehavior: func(r *mock_proto.MockAuthorizationClient) {
				r.EXPECT().SignUp(ctx, mockSignUpInput).Return(
					nil,
					errors.New("some error"),
				)
			},
			expectedResponseBody: `{"status":500,"message":"some error","payload":""}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			authService := mock_proto.NewMockAuthorizationClient(c)
			test.mockBehavior(authService)

			handler := Handler{authMicroservice: authService}

			mux := http.NewServeMux()
			mux.HandleFunc("/api/v1/auth/sign-up", handler.signUp)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/auth/sign-up", bytes.NewBufferString(test.requestBody))
			req = req.WithContext(ctx)
			req.Header.Set("Content-Type", "application/json")

			mux.ServeHTTP(w, req)

			assert.Equal(t, test.expectedResponseBody, w.Body.String())
		})
	}
}

func TestHandler_signIn(t *testing.T) {
	mockSignInInput := &proto.SignInInput{
		Mail:     "test@example.com",
		Password: "password",
	}
	mockSignInInputDto := &dto.SignInInput{
		Mail:     "test@example.com",
		Password: "password",
	}
	mockCookie := &proto.Cookie{Cookie: "mock_cookie"}

	signInInputJSON, _ := json.Marshal(mockSignInInputDto)
	ctx := context.Background()

	tests := []struct {
		name                 string
		requestBody          string
		mockBehavior         func(r *mock_proto.MockAuthorizationClient)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "Sign In Success",
			requestBody: string(signInInputJSON),
			mockBehavior: func(r *mock_proto.MockAuthorizationClient) {
				r.EXPECT().SignIn(ctx, mockSignInInput).Return(mockCookie, nil)
			},
			expectedResponseBody: `{"status":200,"message":"success","payload":""}`,
		},
		{
			name:        "Sign In Error InvalidArgument",
			requestBody: string(signInInputJSON),
			mockBehavior: func(r *mock_proto.MockAuthorizationClient) {
				r.EXPECT().SignIn(ctx, mockSignInInput).Return(
					nil,
					status.Errorf(codes.InvalidArgument, "Неверный аргумент"),
				)
			},
			expectedResponseBody: `{"status":400,"message":"Неверный аргумент","payload":""}`,
		},
		{
			name:        "Sign In InternalServerError",
			requestBody: string(signInInputJSON),
			mockBehavior: func(r *mock_proto.MockAuthorizationClient) {
				r.EXPECT().SignIn(ctx, mockSignInInput).Return(
					nil,
					errors.New("some error"),
				)
			},
			expectedResponseBody: `{"status":500,"message":"some error","payload":""}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			authService := mock_proto.NewMockAuthorizationClient(c)
			test.mockBehavior(authService)

			handler := Handler{authMicroservice: authService}

			mux := http.NewServeMux()
			mux.HandleFunc("/api/v1/auth/login", handler.signIn)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/auth/login", bytes.NewBufferString(test.requestBody))
			req = req.WithContext(ctx)
			req.Header.Set("Content-Type", "application/json")

			mux.ServeHTTP(w, req)

			assert.Equal(t, test.expectedResponseBody, w.Body.String())
		})
	}
}

func TestHandler_logInAdmin(t *testing.T) {
	mockLogInAdminInput := &proto.SignInInput{
		Mail:     "admin@example.com",
		Password: "admin_password",
	}
	mockLogInAdminInputDto := &dto.SignInInput{
		Mail:     "admin@example.com",
		Password: "admin_password",
	}
	mockAdminCookie := &proto.Cookie{Cookie: "mock_admin_cookie"}

	logInAdminInputJSON, _ := json.Marshal(mockLogInAdminInputDto)
	ctx := context.Background()

	tests := []struct {
		name                 string
		requestBody          string
		mockBehavior         func(r *mock_proto.MockAuthorizationClient)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "Log In Admin Success",
			requestBody: string(logInAdminInputJSON),
			mockBehavior: func(r *mock_proto.MockAuthorizationClient) {
				r.EXPECT().LogInAdmin(ctx, mockLogInAdminInput).Return(mockAdminCookie, nil)
			},
			expectedResponseBody: `{"status":200,"message":"success","payload":""}`,
		},
		{
			name:        "Log In Admin Error InvalidArgument",
			requestBody: string(logInAdminInputJSON),
			mockBehavior: func(r *mock_proto.MockAuthorizationClient) {
				r.EXPECT().LogInAdmin(ctx, mockLogInAdminInput).Return(
					nil,
					status.Errorf(codes.InvalidArgument, "Неверный аргумент"),
				)
			},
			expectedResponseBody: `{"status":400,"message":"Неверный аргумент","payload":""}`,
		},
		{
			name:        "Log In Admin InternalServerError",
			requestBody: string(logInAdminInputJSON),
			mockBehavior: func(r *mock_proto.MockAuthorizationClient) {
				r.EXPECT().LogInAdmin(ctx, mockLogInAdminInput).Return(
					nil,
					errors.New("some error"),
				)
			},
			expectedResponseBody: `{"status":500,"message":"some error","payload":""}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			authService := mock_proto.NewMockAuthorizationClient(c)
			test.mockBehavior(authService)

			handler := Handler{authMicroservice: authService}

			mux := http.NewServeMux()
			mux.HandleFunc("/api/v1/auth/admin", handler.logInAdmin)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/auth/admin", bytes.NewBufferString(test.requestBody))
			req = req.WithContext(ctx)
			req.Header.Set("Content-Type", "application/json")

			mux.ServeHTTP(w, req)

			assert.Equal(t, test.expectedResponseBody, w.Body.String())
		})
	}
}
