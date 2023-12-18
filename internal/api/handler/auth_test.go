package handler

// import (
// 	"bytes"
// 	"context"
// 	"errors"
// 	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
// 	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service"
// 	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service/mocks"
// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/assert"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// )

// func TestHandler_signUp(t *testing.T) {
// 	tests := []struct {
// 		name                 string
// 		inputBody            string
// 		inputUser            model.User
// 		mockBehavior         func(r *mock_service.MockAuthorization, user model.User)
// 		expectedStatusCode   int
// 		expectedResponseBody string
// 	}{
// 		{
// 			name:      "Ok",
// 			inputBody: `{"mail": "user@mail.ru", "name": "Test Name", "password": "qwerty"}`,
// 			inputUser: model.User{
// 				Mail:         "user@mail.ru",
// 				Name:         "Test Name",
// 				PasswordHash: "qwerty",
// 			},
// 			mockBehavior: func(r *mock_service.MockAuthorization, user model.User) {
// 				r.EXPECT().CreateUser(gomock.Any(), user).Return(1, nil)
// 				r.EXPECT().GenerateCookie(gomock.Any(), 1).Return("", nil)
// 			},
// 			expectedStatusCode:   http.StatusOK,
// 			expectedResponseBody: `{"status":200,"message":"success","payload":{"id":1}}`,
// 		},
// 		{
// 			name:                 "Wrong input",
// 			inputBody:            `{"mail": "user@mail.ru"}`,
// 			inputUser:            model.User{},
// 			mockBehavior:         func(r *mock_service.MockAuthorization, user model.User) {},
// 			expectedStatusCode:   http.StatusBadRequest,
// 			expectedResponseBody: `{"status":400,"message":"missing required fields","payload":""}`,
// 		},
// 		{
// 			name:                 "Wrong json",
// 			inputBody:            `{"mail": "user@mail.ru"`,
// 			inputUser:            model.User{},
// 			mockBehavior:         func(r *mock_service.MockAuthorization, user model.User) {},
// 			expectedStatusCode:   http.StatusBadRequest,
// 			expectedResponseBody: `{"status":400,"message":"invalid input body","payload":""}`,
// 		},
// 		{
// 			name:      "Already exists email",
// 			inputBody: `{"mail": "user@mail.ru", "name": "Test Name", "password": "qwerty"}`,
// 			inputUser: model.User{
// 				Mail:         "user@mail.ru",
// 				Name:         "Test Name",
// 				PasswordHash: "qwerty",
// 			},
// 			mockBehavior: func(r *mock_service.MockAuthorization, user model.User) {
// 				r.EXPECT().CreateUser(gomock.Any(), user).Return(0, errors.New(""))
// 			},
// 			expectedStatusCode:   http.StatusBadRequest,
// 			expectedResponseBody: `{"status":400,"message":"","payload":""}`,
// 		},
// 		{
// 			name:      "missing required fields",
// 			inputBody: `{"mail": "user@mail.ru", "name": "", "password": "qwerty"}`,
// 			inputUser: model.User{
// 				Mail:         "user@mail.ru",
// 				Name:         "",
// 				PasswordHash: "qwerty",
// 			},
// 			mockBehavior:         func(r *mock_service.MockAuthorization, user model.User) {},
// 			expectedStatusCode:   http.StatusBadRequest,
// 			expectedResponseBody: `{"status":400,"message":"missing required fields","payload":""}`,
// 		},
// 		{
// 			name:      "GenerateCookie error",
// 			inputBody: `{"mail": "user@mail.ru", "name": "Test Name", "password": "qwerty"}`,
// 			inputUser: model.User{
// 				Mail:         "user@mail.ru",
// 				Name:         "Test Name",
// 				PasswordHash: "qwerty",
// 			},
// 			mockBehavior: func(r *mock_service.MockAuthorization, user model.User) {
// 				r.EXPECT().CreateUser(gomock.Any(), user).Return(1, nil)
// 				r.EXPECT().GenerateCookie(gomock.Any(), 1).Return("", errors.New("error"))
// 			},
// 			expectedStatusCode:   http.StatusInternalServerError,
// 			expectedResponseBody: `{"status":500,"message":"error","payload":""}`,
// 		},
// 	}
// 	for _, test := range tests {
// 		t.Serve(test.name, func(t *testing.T) {
// 			c := gomock.NewController(t)
// 			defer c.Finish()

// 			repo := mock_service.NewMockAuthorization(c)
// 			test.mockBehavior(repo, test.inputUser)

// 			services := &service.Service{Authorization: repo}
// 			handler := Handler{services: services}

// 			mux := http.NewServeMux()
// 			mux.HandleFunc("/auth/sign-up", handler.signUp)

// 			w := httptest.NewRecorder()
// 			req := httptest.NewRequest("POST", "/auth/sign-up", bytes.NewBufferString(test.inputBody))

// 			mux.ServeHTTP(w, req)

// 			assert.Equal(t, w.Code, test.expectedStatusCode)
// 			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
// 		})
// 	}
// }

// func TestHandler_signIn(t *testing.T) {
// 	tests := []struct {
// 		name                 string
// 		inputBody            string
// 		inputUser            model.User
// 		mockBehavior         func(r *mock_service.MockAuthorization, user model.User)
// 		expectedStatusCode   int
// 		expectedResponseBody string
// 	}{
// 		{
// 			name:      "Ok",
// 			inputBody: `{"mail": "user@mail.ru", "password": "qwerty"}`,
// 			inputUser: model.User{
// 				Mail:         "user@mail.ru",
// 				PasswordHash: "qwerty",
// 			},
// 			mockBehavior: func(r *mock_service.MockAuthorization, user model.User) {
// 				r.EXPECT().GetUser(gomock.Any(), user.Mail, user.PasswordHash).Return(user, nil)
// 				r.EXPECT().GenerateCookie(gomock.Any(), user.Id).Return("", nil)
// 			},
// 			expectedStatusCode:   200,
// 			expectedResponseBody: `{"status":200,"message":"success","payload":""}`,
// 		},
// 		{
// 			name:                 "Wrong input",
// 			inputBody:            `{"mail": "user@mail.ru"}`,
// 			inputUser:            model.User{},
// 			mockBehavior:         func(r *mock_service.MockAuthorization, user model.User) {},
// 			expectedStatusCode:   400,
// 			expectedResponseBody: `{"status":400,"message":"missing required fields","payload":""}`,
// 		},
// 		{
// 			name:                 "Wrong json",
// 			inputBody:            `{"mail": "user@mail.ru"`,
// 			inputUser:            model.User{},
// 			mockBehavior:         func(r *mock_service.MockAuthorization, user model.User) {},
// 			expectedStatusCode:   400,
// 			expectedResponseBody: `{"status":400,"message":"invalid input body","payload":""}`,
// 		},
// 		{
// 			name:      "invalid mail or password",
// 			inputBody: `{"mail": "user@mail.ru", "password": "qwerty"}`,
// 			inputUser: model.User{
// 				Mail:         "user@mail.ru",
// 				Name:         "Test Name",
// 				PasswordHash: "qwerty",
// 			},
// 			mockBehavior: func(r *mock_service.MockAuthorization, user model.User) {
// 				r.EXPECT().GetUser(gomock.Any(), user.Mail, user.PasswordHash).Return(user, errors.New("error"))
// 			},
// 			expectedStatusCode:   401,
// 			expectedResponseBody: `{"status":401,"message":"invalid mail or password","payload":""}`,
// 		},
// 		{
// 			name:      "GenerateCookie error",
// 			inputBody: `{"mail": "user@mail.ru", "name": "Test Name", "password": "qwerty"}`,
// 			inputUser: model.User{
// 				Mail:         "user@mail.ru",
// 				Name:         "Test Name",
// 				PasswordHash: "qwerty",
// 			},
// 			mockBehavior: func(r *mock_service.MockAuthorization, user model.User) {
// 				r.EXPECT().GetUser(gomock.Any(), user.Mail, user.PasswordHash).Return(user, nil)
// 				r.EXPECT().GenerateCookie(gomock.Any(), user.Id).Return("", errors.New("error"))
// 			},
// 			expectedStatusCode:   http.StatusInternalServerError,
// 			expectedResponseBody: `{"status":500,"message":"error","payload":""}`,
// 		},
// 	}
// 	for _, test := range tests {
// 		t.Serve(test.name, func(t *testing.T) {
// 			c := gomock.NewController(t)
// 			defer c.Finish()

// 			repo := mock_service.NewMockAuthorization(c)
// 			test.mockBehavior(repo, test.inputUser)

// 			ctx := context.Background()
// 			services := &service.Service{Authorization: repo}
// 			handler := Handler{services, ctx}

// 			mux := http.NewServeMux()
// 			mux.HandleFunc("/auth/login", handler.signIn)

// 			w := httptest.NewRecorder()
// 			req := httptest.NewRequest("POST", "/auth/login", bytes.NewBufferString(test.inputBody))

// 			mux.ServeHTTP(w, req)

// 			assert.Equal(t, w.Code, test.expectedStatusCode)
// 			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
// 		})
// 	}
// }

// func TestHandler_logout(t *testing.T) {
// 	mockCookie := &http.Cookie{
// 		Name:  "session_id",
// 		Value: "value",
// 	}
// 	tests := []struct {
// 		name                 string
// 		inputCookie          *http.Cookie
// 		mockBehavior         func(r *mock_service.MockAuthorization, cookie *http.Cookie)
// 		expectedStatusCode   int
// 		expectedResponseBody string
// 	}{
// 		{
// 			name:        "Ok",
// 			inputCookie: mockCookie,
// 			mockBehavior: func(r *mock_service.MockAuthorization, cookie *http.Cookie) {
// 				r.EXPECT().DeleteCookie(gomock.Any(), cookie.Value).Return(nil)
// 			},
// 			expectedStatusCode:   200,
// 			expectedResponseBody: `{"status":200,"message":"success","payload":""}`,
// 		},
// 		{
// 			name:        "invalid mail or password",
// 			inputCookie: mockCookie,
// 			mockBehavior: func(r *mock_service.MockAuthorization, cookie *http.Cookie) {
// 				r.EXPECT().DeleteCookie(gomock.Any(), cookie.Value).Return(errors.New("error"))
// 			},
// 			expectedStatusCode:   500,
// 			expectedResponseBody: `{"status":500,"message":"Invalid cookie deletion","payload":""}`,
// 		},
// 		{
// 			name:                 "No cookie",
// 			inputCookie:          nil,
// 			mockBehavior:         func(r *mock_service.MockAuthorization, cookie *http.Cookie) {},
// 			expectedStatusCode:   401,
// 			expectedResponseBody: `{"status":401,"message":"no session","payload":""}`,
// 		},
// 	}
// 	for _, test := range tests {
// 		t.Serve(test.name, func(t *testing.T) {
// 			c := gomock.NewController(t)
// 			defer c.Finish()

// 			repo := mock_service.NewMockAuthorization(c)
// 			cookie := &http.Cookie{
// 				Name:  "session_id",
// 				Value: "value",
// 			}
// 			test.mockBehavior(repo, cookie)

// 			ctx := context.Background()
// 			services := &service.Service{Authorization: repo}
// 			handler := Handler{services, ctx}

// 			mux := http.NewServeMux()
// 			mux.HandleFunc("/auth/logout", handler.logout)

// 			w := httptest.NewRecorder()
// 			req := httptest.NewRequest("GET", "/auth/logout", nil)

// 			if test.inputCookie != nil {
// 				req.AddCookie(test.inputCookie)
// 			}

// 			mux.ServeHTTP(w, req)

// 			assert.Equal(t, w.Code, test.expectedStatusCode)
// 			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
// 		})
// 	}
// }
