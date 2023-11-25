package handler

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/admin/proto"
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
	var stat model.Statistic
	if err := decoder.Decode(&stat); err != nil {
		newErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid input body")
		return
	}
	_, err := h.adminMicroservice.CreateStatistic(
		r.Context(),
		&proto.Statistic{
			UserId: int32(r.Context().Value(keyUserID).(int)),
			Rating: int32(*stat.Rating),
			//Liked: stat.Liked,

		})

	if err != nil {
		statusCode, message := parseError(err)
		newErrorClientResponseDto(r.Context(), w, statusCode, message)
		return
	}
	NewSuccessClientResponseDto(r.Context(), w, "")
}
