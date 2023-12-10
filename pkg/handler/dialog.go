package handler

import (
	"github.com/go-park-mail-ru/2023_2_Umlaut/static"
	"github.com/gorilla/mux"
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
	userId := r.Context().Value(static.KeyUserID).(int)

	dialogs, err := h.services.Dialog.GetDialogs(r.Context(), userId)
	if err != nil {
		newErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}

	//err = h.addDialogsToUserHub(w, r, userId, dialogs)
	//if err != nil {
	//	newErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, err.Error())
	//	return
	//}

	NewSuccessClientResponseArrayDto(r.Context(), w, dialogs)
}

// @Summary get dialog message
// @Tags dialog
// @Accept  json
// @Param id path integer true "Recipient ID"
// @Produce  json
// @Success 200 {object} ClientResponseDto[[]model.Message]
// @Failure 401,500 {object} ClientResponseDto[string]
// @Router /api/v1/dialogs/{id}/message [get]
func (h *Handler) getDialogMessage(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(static.KeyUserID).(int)
	recipientId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		newErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid params")
		return
	}

	messages, err := h.services.Message.GetDialogMessages(r.Context(), userId, recipientId)
	if err != nil {
		newErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	} else if len(messages) > 0 && *messages[0].RecipientId != userId && *messages[0].SenderId != userId {
		newErrorClientResponseDto(r.Context(), w, http.StatusForbidden, "нет доступа")
	}

	NewSuccessClientResponseArrayDto(r.Context(), w, messages)
}
