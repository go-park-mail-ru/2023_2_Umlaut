package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
)

// @Summary create user like
// @Tags like
// @ID like
// @Accept  json
// @Produce  json
// @Param input body likeDto true "Like data to update"
// @Success 200 {object} ClientResponseDto[string]
// @Failure 400,401,404 {object} ClientResponseDto[string]
// @Router /api/v1/like [post]
func (h *Handler) createLike(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var likeJson likeDto
	if err := decoder.Decode(&likeJson); err != nil {
		newErrorClientResponseDto(w, http.StatusBadRequest, err.Error())
		return
	}

	userId := r.Context().Value(keyUserID).(int)
	like := model.Like{LikedByUserId: userId, LikedToUserId: likeJson.LikedUserId, CommittedAt: time.Time(likeJson.CommittedAt)}

	if err := h.services.CreateLike(r.Context(), like); err != nil {
		newErrorClientResponseDto(w, http.StatusBadRequest, err.Error())
		return
	}

	exists, err := h.services.IsUserLiked(r.Context(), like)
	if err != nil {
		newErrorClientResponseDto(w, http.StatusInternalServerError, err.Error())
		return
	}

	if exists {
		NewSuccessClientResponseDto(w, "Matching likes")
		return
	}
	NewSuccessClientResponseDto(w, "")
}
