package handler

import (
	"errors"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model/ws"
	"io"
	"net/http"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/static"
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
	body, err := io.ReadAll(r.Body)
	if err != nil {
		newErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid input body")
		return
	}
	var like model.Like
	if err := like.UnmarshalJSON(body); err != nil {
		newErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid input body")
		return
	}

	userId := r.Context().Value(static.KeyUserID).(int)
	like.LikedByUserId = userId

	dialog, err := h.services.Like.CreateLike(r.Context(), like)
	if err != nil {
		if errors.Is(err, static.ErrAlreadyExists) {
			newErrorClientResponseDto(r.Context(), w, http.StatusOK, "already liked")
			return
		}
		if errors.Is(err, static.ErrMutualLike) {

			h.hub.Broadcast <- &ws.Notification{
				Type:    static.Match,
				Payload: dialog,
			}
			newErrorClientResponseDto(r.Context(), w, http.StatusOK, "Mutual like")
			return
		}
		newErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}
	NewSuccessClientResponseDto(r.Context(), w, "")
}
