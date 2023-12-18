package handler

import (
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/constants"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/core/chat"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/dto"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// @Summary register user to WebSocket hub
// @Description Registers a user to the WebSocket hub and initiates connection
// @Tags websocket
// @ID registerUserToHub
// @Accept json
// @Produce json
// @Success 200 {object} ClientResponseDto[string]
// @Failure 500 {object} ClientResponseDto[string]
// @Router /api/v1/ws/messenger [get]
func (h *Handler) registerUserToHub(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(constants.KeyUserID).(int)
	if _, ok := h.hub.Users[userId]; ok {
		return
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, "can not open connect")
		return
	}
	cl := &chat.Client{
		Id:            userId,
		Notifications: make(chan *chat.Notification, 10),
		Conn:          conn,
		Logger:        h.logger,
	}
	h.hub.Register <- cl

	go cl.WriteMessage()
	cl.ReadMessage(r.Context(), h.hub, h.services)
}
