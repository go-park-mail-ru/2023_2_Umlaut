package handler

import (
	"encoding/json"
	"errors"
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
// @Failure 500 {object} ClientResponseDto[string]
// @Router /api/v1/like [post]
func (h *Handler) createLike(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var like model.Like
	if err := decoder.Decode(&like); err != nil {
		newErrorClientResponseDto(h.ctx, w, http.StatusBadRequest, "invalid input body")
		return
	}
	userId := r.Context().Value(keyUserID).(int)
	like.LikedByUserId = userId

	err := h.services.Like.CreateLike(r.Context(), like)
	if err != nil {
		if errors.Is(err, model.AlreadyExists) {
			newErrorClientResponseDto(h.ctx, w, http.StatusOK, "already liked")
			return
		}
		if errors.Is(err, model.MutualLike) {
			newErrorClientResponseDto(h.ctx, w, http.StatusOK, "Mutual like")
			return
		}
		newErrorClientResponseDto(h.ctx, w, http.StatusInternalServerError, err.Error())
		return
	}
	NewSuccessClientResponseDto(h.ctx, w, "")
}
