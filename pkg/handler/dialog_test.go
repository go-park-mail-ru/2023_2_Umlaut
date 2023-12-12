package handler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/service/mocks"
	"github.com/go-park-mail-ru/2023_2_Umlaut/static"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	// gmux "github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_getDialogs(t *testing.T) {
	mockUserID := 1
	mockDialogs := []model.Dialog{
		{Id: 1, User1Id: mockUserID, User2Id: 2},
		{Id: 2, User1Id: mockUserID, User2Id: 3},
	}
	response := ClientResponseArrayDto[model.Dialog]{
		Status:  200,
		Message: "success",
		Payload: mockDialogs,
	}

	dialogsJSON, _ := json.Marshal(response)

	tests := []struct {
		name                 string
		mockBehavior         func(r *mock_service.MockDialog)
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Success",
			mockBehavior: func(r *mock_service.MockDialog) {
				r.EXPECT().GetDialogs(gomock.Any(), mockUserID).Return(mockDialogs, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: string(dialogsJSON),
		},
		{
			name: "Error in GetDialogs",
			mockBehavior: func(r *mock_service.MockDialog) {
				r.EXPECT().GetDialogs(gomock.Any(), mockUserID).Return(nil, errors.New("some error"))
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"status":500,"message":"some error","payload":""}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repoDialog := mock_service.NewMockDialog(c)
			test.mockBehavior(repoDialog)

			ctx := context.WithValue(context.Background(), static.KeyUserID, mockUserID)
			services := &service.Service{Dialog: repoDialog}
			handler := Handler{services: services}

			mux := http.NewServeMux()
			mux.HandleFunc("/api/v1/dialogs", handler.getDialogs)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/api/v1/dialogs", nil)
			mux.ServeHTTP(w, req.WithContext(ctx))

			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
