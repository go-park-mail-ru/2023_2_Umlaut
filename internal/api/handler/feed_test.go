package handler

// import (
// 	"context"
// 	"encoding/json"
// 	"errors"
// 	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
// 	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service"
// 	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service/mocks"
// 	"github.com/go-park-mail-ru/2023_2_Umlaut/static"
// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/assert"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// )

// func TestHandler_feed(t *testing.T) {
// 	mockCookie := &http.Cookie{
// 		Name:  "session_id",
// 		Value: "value",
// 	}
// 	mockUser := model.User{Mail: "max@max.ru", PasswordHash: "passWord", Name: "Max"}
// 	response := ClientResponseDto[model.User]{
// 		Status:  200,
// 		Message: "success",
// 		Payload: mockUser,
// 	}
// 	jsonData, _ := json.Marshal(response)

// 	tests := []struct {
// 		name                 string
// 		inputCookie          *http.Cookie
// 		mockBehavior         func(r *mock_service.MockFeed)
// 		expectedStatusCode   int
// 		expectedResponseBody string
// 	}{
// 		{
// 			name:        "Ok",
// 			inputCookie: mockCookie,
// 			mockBehavior: func(r *mock_service.MockFeed) {
// 				r.EXPECT().GetNextUser(gomock.Any(), 1).Return(mockUser, nil)
// 			},
// 			expectedStatusCode:   http.StatusOK,
// 			expectedResponseBody: string(jsonData),
// 		},
// 		{
// 			name:        "Error",
// 			inputCookie: mockCookie,
// 			mockBehavior: func(r *mock_service.MockFeed) {
// 				r.EXPECT().GetNextUser(gomock.Any(), 1).Return(mockUser, errors.New("some error"))
// 			},
// 			expectedStatusCode:   http.StatusInternalServerError,
// 			expectedResponseBody: `{"status":500,"message":"some error","payload":""}`,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Serve(test.name, func(t *testing.T) {
// 			c := gomock.NewController(t)
// 			defer c.Finish()

// 			repo := mock_service.NewMockFeed(c)
// 			test.mockBehavior(repo)

// 			ctx := context.WithValue(context.Background(), static.KeyUserID, 1)
// 			services := &service.Service{Feed: repo}
// 			handler := Handler{services, ctx}

// 			mux := http.NewServeMux()
// 			mux.HandleFunc("/api/v1/feed", handler.feed)

// 			w := httptest.NewRecorder()
// 			req := httptest.NewRequest("GET", "/api/v1/feed", nil)
// 			if test.inputCookie != nil {
// 				req.AddCookie(test.inputCookie)
// 			}

// 			mux.ServeHTTP(w, req.WithContext(ctx))

// 			assert.Equal(t, w.Code, test.expectedStatusCode)
// 			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
// 		})
// 	}
// }
