package handler

import (
	"net/http"
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

	NewSuccessClientResponseArrayDto(r.Context(), w, dialogs)
}

// @Summary get dialog message
// @Tags dialog
// @ID dialogMsg
// @Accept  json
// @Produce  json
// @Success 200 {object} ClientResponseDto[[]model.Dialog]
// @Failure 401,500 {object} ClientResponseDto[string]
// @Router /api/v1/dialog [get]
func (h *Handler) getDialogMessage(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(keyUserID).(int)

	dialogs, err := h.services.Dialog.GetDialogs(r.Context(), userId)
	if err != nil {
		newErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessClientResponseArrayDto(r.Context(), w, dialogs)
}
