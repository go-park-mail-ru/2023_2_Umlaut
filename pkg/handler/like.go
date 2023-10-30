package handler

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"net/http"
)

// @Summary create user like
// @Tags like
// @ID like
// @Accept  json
// @Produce  json
// @Param input body model.Like true "Like data to update"
// @Success 200 {object} ClientResponseDto[string]
// @Failure 400,401,404 {object} ClientResponseDto[string]
// @Router /api/v1/like [post]
func (h *Handler) createLike(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var like model.Like
	if err := decoder.Decode(&like); err != nil {
		newErrorClientResponseDto(h.ctx, w, http.StatusBadRequest, err.Error())
		return
	}

	userId := r.Context().Value(keyUserID).(int)
	like.LikedByUserId = userId

	if err := h.services.CreateLike(r.Context(), like); err != nil {
		newErrorClientResponseDto(h.ctx, w, http.StatusBadRequest, err.Error())
		return
	}

	exists, err := h.services.IsUserLiked(r.Context(), like)
	if err != nil {
		newErrorClientResponseDto(h.ctx, w, http.StatusInternalServerError, err.Error())
		return
	}

	if exists {
		NewSuccessClientResponseDto(h.ctx, w, "Matching likes")
		return
	}
	NewSuccessClientResponseDto(h.ctx, w, "")
}
