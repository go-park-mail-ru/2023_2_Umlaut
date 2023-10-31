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

	dialogs, err := h.services.GetDialogs(r.Context(), userId)
	if err != nil {
		newErrorClientResponseDto(h.ctx, w, http.StatusInternalServerError, err.Error())
	}

	NewSuccessClientResponseArrayDto(h.ctx, w, dialogs)
}
