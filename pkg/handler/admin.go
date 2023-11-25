package handler

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/admin/proto"
	"github.com/go-park-mail-ru/2023_2_Umlaut/utils"
	"net/http"
)

// @Summary create statistic
// @Tags statistic
// @ID statistic
// @Accept  json
// @Produce  json
// @Param input body model.Statistic true "Statistic data"
// @Success 200 {object} ClientResponseDto[string]
// @Failure 500 {object} ClientResponseDto[string]
// @Router /api/v1/like [post]
func (h *Handler) CreateStatistic(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var stat model.Feedback
	if err := decoder.Decode(&stat); err != nil {
		newErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid input body")
		return
	}
	_, err := h.adminMicroservice.CreateStatistic(
		r.Context(),
		&proto.Feedback{
			UserId:     int32(r.Context().Value(keyUserID).(int)),
			Rating:     utils.ModifyInt(stat.Rating),
			Liked:      utils.ModifyString(stat.Liked),
			NeedFix:    utils.ModifyString(stat.NeedFix),
			CommentFix: utils.ModifyString(stat.CommentFix),
			Comment:    utils.ModifyString(stat.Comment),
			Show:       stat.Show,
		})

	if err != nil {
		statusCode, message := parseError(err)
		newErrorClientResponseDto(r.Context(), w, statusCode, message)
		return
	}
	NewSuccessClientResponseDto(r.Context(), w, "")
}
