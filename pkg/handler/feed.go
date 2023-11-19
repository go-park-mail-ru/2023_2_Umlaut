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
		newErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}
	NewSuccessClientResponseDto(r.Context(), w, nextUser)
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
		newErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}
	if nextUsers == nil {
		newErrorClientResponseDto(r.Context(), w, http.StatusNotFound, "out of users")
		return
	}

	NewSuccessClientResponseArrayDto(r.Context(), w, nextUsers)
}
