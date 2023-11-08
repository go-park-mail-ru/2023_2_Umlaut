package handler

import (
	"net/http"
)

// @Summary get user for feed
// @Tags feed
// @ID feed
// @Accept  json
// @Produce  json
// @Success 200 {object} ClientResponseDto[model.User]
// @Failure 500 {object} ClientResponseDto[string]
// @Router /api/v1/feed [get]
func (h *Handler) feed(w http.ResponseWriter, r *http.Request) {
	nextUser, err := h.services.Feed.GetNextUser(r.Context(), r.Context().Value(keyUserID).(int))
	if err != nil {
		newErrorClientResponseDto(h.ctx, w, http.StatusInternalServerError, err.Error())
		return
	}
	NewSuccessClientResponseDto(h.ctx, w, nextUser)
}

// @Summary get users for feed
// @Tags feed
// @Accept  json
// @Produce  json
// @Success 200 {object} ClientResponseDto[[]model.User]
// @Failure 500 {object} ClientResponseDto[string]
// @Router /api/v1/feed/users [get]
func (h *Handler) getNextUsers(w http.ResponseWriter, r *http.Request) {
	nextUsers, err := h.services.Feed.GetNextUsers(r.Context(), r.Context().Value(keyUserID).(int))
	if err != nil {
		newErrorClientResponseDto(h.ctx, w, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessClientResponseArrayDto(h.ctx, w, nextUsers)
}
