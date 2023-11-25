package handler

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/admin/proto"
	"github.com/go-park-mail-ru/2023_2_Umlaut/utils"
	"math/rand"
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
// @Router /api/v1/feedback [post]
func (h *Handler) createFeedback(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var stat model.Feedback
	if err := decoder.Decode(&stat); err != nil {
		newErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid input body")
		return
	}
	_, err := h.adminMicroservice.CreateFeedback(
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

// @Summary create recommendation
// @Tags statistic
// @ID statistic
// @Accept  json
// @Produce  json
// @Param input body model.Recommendation true "Recommendation data"
// @Success 200 {object} ClientResponseDto[string]
// @Failure 500 {object} ClientResponseDto[string]
// @Router /api/v1/recommendation [post]
func (h *Handler) createRecommendation(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var rec model.Recommendation
	if err := decoder.Decode(&rec); err != nil {
		newErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid input body")
		return
	}
	_, err := h.adminMicroservice.CreateRecommendation(
		r.Context(),
		&proto.Recommendation{
			UserId:    int32(r.Context().Value(keyUserID).(int)),
			Recommend: utils.ModifyInt(rec.Recommend),
			Show:      rec.Show,
		})

	if err != nil {
		statusCode, message := parseError(err)
		newErrorClientResponseDto(r.Context(), w, statusCode, message)
		return
	}
	NewSuccessClientResponseDto(r.Context(), w, "")
}

// @Summary create feed feedback
// @Tags statistic
// @ID statistic
// @Accept  json
// @Produce  json
// @Param input body model.Recommendation true "feed_feedback data"
// @Success 200 {object} ClientResponseDto[string]
// @Failure 500 {object} ClientResponseDto[string]
// @Router /api/v1/recommendation [post]
func (h *Handler) createFeedFeedback(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var rec model.Recommendation
	if err := decoder.Decode(&rec); err != nil {
		newErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid input body")
		return
	}
	_, err := h.adminMicroservice.CreateRecommendation(
		r.Context(),
		&proto.Recommendation{
			UserId:    int32(r.Context().Value(keyUserID).(int)),
			Recommend: utils.ModifyInt(rec.Recommend),
			Show:      rec.Show,
		})

	if err != nil {
		statusCode, message := parseError(err)
		newErrorClientResponseDto(r.Context(), w, statusCode, message)
		return
	}
	NewSuccessClientResponseDto(r.Context(), w, "")
}

// @Summary show csat for user
// @Tags statistic
// @ID statistic
// @Accept  json
// @Produce  json
// @Success 200 {object} ClientResponseDto[int]
// @Failure 500 {object} ClientResponseDto[string]
// @Router /api/v1/show-csat [get]
func (h *Handler) showCSAT(w http.ResponseWriter, r *http.Request) {
	NewSuccessClientResponseDto(r.Context(), w, rand.Intn(4))
}

// @Summary statistic by recommendation
// @Tags statistic
// @ID statistic
// @Accept  json
// @Produce  json
// @Success 200 {object} ClientResponseDto[RecommendationStatistic]
// @Failure 500 {object} ClientResponseDto[string]
// @Router /api/v1/show-csat [get]
func (h *Handler) recommendationStatistic(w http.ResponseWriter, r *http.Request) {
	NewSuccessClientResponseDto(r.Context(), w, rand.Intn(4))
}
