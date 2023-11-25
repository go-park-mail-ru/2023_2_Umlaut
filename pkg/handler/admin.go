package handler

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/static"
	"net/http"
)

// CreateStatistic @Summary create user like
// @Tags like
// @ID like
// @Accept  json
// @Produce  json
// @Param input body model.Like true "Like data to update"
// @Success 200 {object} ClientResponseDto[string]
// @Failure 500 {object} ClientResponseDto[string]
// @Router /api/v1/like [post]
func (h *Handler) CreateStatistic(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var like model.Statistic
	if err := decoder.Decode(&like); err != nil {
		newErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid input body")
		return
	}
	userId := r.Context().Value(keyUserID).(int)
	like.LikedByUserId = userId

	err := h.services.Like.CreateLike(r.Context(), like)
	if err != nil {
		if errors.Is(err, static.ErrAlreadyExists) {
			newErrorClientResponseDto(r.Context(), w, http.StatusOK, "already liked")
			return
		}
		if errors.Is(err, static.ErrMutualLike) {
			newErrorClientResponseDto(r.Context(), w, http.StatusOK, "Mutual like")
			return
		}
		newErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}
	NewSuccessClientResponseDto(r.Context(), w, "")
}
