package handler

import (
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"net/http"
)

// @Summary get user for feed
// @Tags feed
// @ID feed
// @Accept  json
// @Produce  json
// @Success 200 {object} ClientResponseDto[model.User]
// @Failure 401,404 {object} ClientResponseDto[string]
// @Router /api/v1/feed [get]
func (h *Handler) feed(w http.ResponseWriter, r *http.Request) {
	nextUser, err := h.services.GetNextUser(r.Context(), r.Context().Value(keyUserID).(int))
	if err != nil {
		newErrorClientResponseDto(w, http.StatusInternalServerError, err.Error())
		return
	}
	NewSuccessClientResponseDto[model.User](w, nextUser)
}
