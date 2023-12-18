package handler

import (
	"errors"
	static2 "github.com/go-park-mail-ru/2023_2_Umlaut/pkg/constants"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/core"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/core/chat"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/dto"
	"io"
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
	body, err := io.ReadAll(r.Body)
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid input body")
		return
	}
	var like core.Like
	if err := like.UnmarshalJSON(body); err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid input body")
		return
	}

	userId := r.Context().Value(static2.KeyUserID).(int)
	like.LikedByUserId = userId

	dialog, err := h.services.Like.CreateLike(r.Context(), like)
	if err != nil {
		if errors.Is(err, static2.ErrAlreadyExists) {
			dto.NewErrorClientResponseDto(r.Context(), w, http.StatusOK, "already liked")
			return
		}
		if errors.Is(err, static2.ErrMutualLike) {

			h.hub.Broadcast <- &chat.Notification{
				Type:    static2.Match,
				Payload: dialog,
			}
			dto.NewErrorClientResponseDto(r.Context(), w, http.StatusOK, "Mutual like")
			return
		}
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}
	dto.NewSuccessClientResponseDto(r.Context(), w, "")
}

// @Summary get users who have liked the user
// @Tags like
// @Accept  json
// @Produce  json
// @Success 200 {object} ClientResponseDto[model.PremiumLike]
// @Failure 401,402,403,500 {object} ClientResponseDto[string]
// @Router /api/v1/premium/likes [get]
func (h *Handler) getPremiumLikes(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(static2.KeyUserID).(int)

	show, likes, err := h.services.Like.GetUserLikedToLikes(r.Context(), userId)
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}
	dto.NewSuccessClientResponseDto(r.Context(), w, map[string]interface{}{
		"likes": likes,
		"show":  show,
	})
}
