package handler

import (
	"bytes"
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

func TestHandler_signUp(t *testing.T) {
	tests := []struct {
		name                 string
		inputBody            string
		inputUser            model.User
		mockBehavior         func(r *mock_service.MockAuthorization, user model.User)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"mail": "user@mail.ru", "name": "Test Name", "password": "qwerty"}`,
			inputUser: model.User{
				Mail:         "user@mail.ru",
				Name:         "Test Name",
				PasswordHash: "qwerty",
			},
			mockBehavior: func(r *mock_service.MockAuthorization, user model.User) {
				r.EXPECT().CreateUser(user).Return(1, nil)
				r.EXPECT().GenerateCookie(gomock.Any(), user.Id).Return("", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:                 "Wrong input",
			inputBody:            `{"mail": "user@mail.ru"}`,
			inputUser:            model.User{},
			mockBehavior:         func(r *mock_service.MockAuthorization, user model.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalidate fields"}`,
		},
		{
			name:                 "Wrong json",
			inputBody:            `{"mail": "user@mail.ru"`,
			inputUser:            model.User{},
			mockBehavior:         func(r *mock_service.MockAuthorization, user model.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Already exists email",
			inputBody: `{"mail": "user@mail.ru", "name": "Test Name", "password": "qwerty"}`,
			inputUser: model.User{
				Mail:         "user@mail.ru",
				Name:         "Test Name",
				PasswordHash: "qwerty",
			},
			mockBehavior: func(r *mock_service.MockAuthorization, user model.User) {
				r.EXPECT().CreateUser(user).Return(0, errors.New(""))
			},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"Account with this email already exists"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockAuthorization(c)
			test.mockBehavior(repo, test.inputUser)

			services := &service.Service{Authorization: repo}
			handler := Handler{services}

			mux := http.NewServeMux()
			mux.HandleFunc("/auth/sign-up", handler.signUp)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/auth/sign-up", bytes.NewBufferString(test.inputBody))

			mux.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
