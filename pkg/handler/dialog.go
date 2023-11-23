package handler

import (
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/handler/ws"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
)

// @Summary get user dialogs
// @Tags dialog
// @ID dialog
// @Accept  json
// @Produce  json
// @Success 200 {object} ClientResponseDto[[]model.Dialog]
// @Failure 401,500 {object} ClientResponseDto[string]
// @Router /api/v1/dialogs [get]
func (h *Handler) getDialogs(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(keyUserID).(int)

	dialogs, err := h.services.Dialog.GetDialogs(r.Context(), userId)
	if err != nil {
		newErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}

	for _, dialog := range dialogs {
		err = h.addDialogToHub(w, r, dialog, userId)
		if err != nil {
			newErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, err.Error())
			return
		}
	}

	NewSuccessClientResponseArrayDto(r.Context(), w, dialogs)
}

// @Summary get dialog message
// @Tags dialog
// @Accept  json
// @Param id path integer true "Dialog ID"
// @Produce  json
// @Success 200 {object} ClientResponseDto[[]model.Message]
// @Failure 401,500 {object} ClientResponseDto[string]
// @Router /api/v1/dialogs/{id}/message [get]
func (h *Handler) getDialogMessage(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		newErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid params")
		return
	}

	messages, err := h.services.Dialog.GetDialogMessages(r.Context(), id)
	if err != nil {
		newErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessClientResponseArrayDto(r.Context(), w, messages)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) addDialogToHub(w http.ResponseWriter, r *http.Request, dialog model.Dialog, userId int) error {
	h.hub.Dialogs[dialog.Id] = &ws.Dialog{
		Id:                  dialog.Id,
		User1Id:             dialog.User1Id,
		User2Id:             dialog.User2Id,
		Сompanion:           dialog.Сompanion,
		СompanionImagePaths: dialog.СompanionImagePaths,
		LastMessage:         dialog.LastMessage,
		Clients:             make(map[int]*ws.Client),
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	cl := &ws.Client{
		Conn:     conn,
		Message:  make(chan *model.Message, 10),
		Id:       userId,
		DialogId: dialog.Id,
	}
	text := "HELLO FROM WS"
	m := &model.Message{
		Text:     &text,
		DialogId: &dialog.Id,
		SenderId: &userId,
	}
	h.hub.Register <- cl
	h.hub.Broadcast <- m

	go cl.WriteMessage()
	cl.ReadMessage(h.hub)
	return nil
}
