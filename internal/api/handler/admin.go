package handler

import (
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/constants"
	core2 "github.com/go-park-mail-ru/2023_2_Umlaut/internal/model/core"
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/model/dto"
	utils2 "github.com/go-park-mail-ru/2023_2_Umlaut/internal/utils"
	"io"
	"net/http"

	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/microservices/admin/proto"
)

// @Summary create statistic
// @Tags statistic
// @ID Feedback
// @Accept  json
// @Produce  json
// @Param input body model.Feedback true "Statistic data"
// @Success 200 {object} ClientResponseDto[string]
// @Failure 500 {object} ClientResponseDto[string]
// @Router /api/v1/feedback [post]
func (h *Handler) createFeedback(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid input body")
		return
	}
	var stat core2.Feedback
	if err := stat.UnmarshalJSON(body); err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid input body")
		return
	}
	_, err = h.adminMicroservice.CreateFeedback(
		r.Context(),
		&proto.Feedback{
			UserId:  int32(r.Context().Value(constants.KeyUserID).(int)),
			Rating:  utils2.ModifyInt(stat.Rating),
			Liked:   utils2.ModifyString(stat.Liked),
			NeedFix: utils2.ModifyString(stat.NeedFix),
			Comment: utils2.ModifyString(stat.Comment),
		})

	if err != nil {
		statusCode, message := utils2.ParseError(err)
		dto.NewErrorClientResponseDto(r.Context(), w, statusCode, message)
		return
	}
	dto.NewSuccessClientResponseDto(r.Context(), w, "")
}

// @Summary create recommendation
// @Tags statistic
// @ID Recommendation
// @Accept  json
// @Produce  json
// @Param input body model.Recommendation true "Recommendation data"
// @Success 200 {object} ClientResponseDto[string]
// @Failure 500 {object} ClientResponseDto[string]
// @Router /api/v1/recommendation [post]
func (h *Handler) createRecommendation(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid input body")
		return
	}
	var rec core2.Recommendation
	if err := rec.UnmarshalJSON(body); err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid input body")
		return
	}
	_, err = h.adminMicroservice.CreateRecommendation(
		r.Context(),
		&proto.Recommendation{
			UserId: int32(r.Context().Value(constants.KeyUserID).(int)),
			Rating: utils2.ModifyInt(rec.Rating),
		})
	if err != nil {
		statusCode, message := utils2.ParseError(err)
		dto.NewErrorClientResponseDto(r.Context(), w, statusCode, message)
		return
	}
	dto.NewSuccessClientResponseDto(r.Context(), w, "")
}

// @Summary create feed feedback
// @Tags statistic
// @ID FeedFeedback
// @Accept  json
// @Produce  json
// @Param input body model.Recommendation true "feed_feedback data"
// @Success 200 {object} ClientResponseDto[string]
// @Failure 500 {object} ClientResponseDto[string]
// @Router /api/v1/feed-feedback [post]
func (h *Handler) createFeedFeedback(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid input body")
		return
	}
	var rec core2.Recommendation
	if err := rec.UnmarshalJSON(body); err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusBadRequest, "invalid input body")
		return
	}
	_, err = h.adminMicroservice.CreateRecommendation(
		r.Context(),
		&proto.Recommendation{
			UserId: int32(r.Context().Value(constants.KeyUserID).(int)),
			Rating: utils2.ModifyInt(rec.Rating),
		})

	if err != nil {
		statusCode, message := utils2.ParseError(err)
		dto.NewErrorClientResponseDto(r.Context(), w, statusCode, message)
		return
	}
	dto.NewSuccessClientResponseDto(r.Context(), w, "")
}

// @Summary show csat for user
// @Tags statistic
// @ID CSAT
// @Accept  json
// @Produce  json
// @Success 200 {object} ClientResponseDto[int]
// @Failure 500 {object} ClientResponseDto[string]
// @Router /api/v1/show-csat [get]
func (h *Handler) showCSAT(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(constants.KeyUserID).(int)

	csatType, err := h.services.Admin.GetCSATType(r.Context(), id)
	if err != nil {
		dto.NewErrorClientResponseDto(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}
	dto.NewSuccessClientResponseDto(r.Context(), w, csatType)
}

// @Summary statistic by recommendation
// @Tags statistic
// @ID RecommendationStatistic
// @Accept  json
// @Produce  json
// @Success 200 {object} ClientResponseDto[model.RecommendationStatistic]
// @Failure 500 {object} ClientResponseDto[string]
// @Router /api/v1/admin/recommendation [get]
func (h *Handler) getRecommendationStatistic(w http.ResponseWriter, r *http.Request) {
	recommend, err := h.adminMicroservice.GetRecommendationStatistic(r.Context(), &proto.AdminEmpty{})
	if err != nil {
		statusCode, message := utils2.ParseError(err)
		dto.NewErrorClientResponseDto(r.Context(), w, statusCode, message)
		return
	}

	dto.NewSuccessClientResponseDto(r.Context(), w, &core2.RecommendationStatistic{
		AvgRecommend:   recommend.AvgRecommend,
		NPS:            recommend.NPS,
		RecommendCount: recommend.RecommendCount,
	})
}

// @Summary statistic by feedback
// @Tags statistic
// @ID FeedbackStatistic
// @Accept  json
// @Produce  json
// @Success 200 {object} ClientResponseDto[model.FeedbackStatistic]
// @Failure 500 {object} ClientResponseDto[string]
// @Router /api/v1/admin/feedback [get]
func (h *Handler) getFeedbackStatistic(w http.ResponseWriter, r *http.Request) {
	feedbackStat, err := h.adminMicroservice.GetFeedbackStatistic(r.Context(), &proto.AdminEmpty{})
	if err != nil {
		statusCode, message := utils2.ParseError(err)
		dto.NewErrorClientResponseDto(r.Context(), w, statusCode, message)
		return
	}

	dto.NewSuccessClientResponseDto(r.Context(), w, &core2.FeedbackStatistic{
		AvgRating:   feedbackStat.AvgRating,
		RatingCount: feedbackStat.RatingCount,
		LikedMap:    utils2.ModifyLikedMap(feedbackStat.LikedMap),
		NeedFixMap:  utils2.ModifyNeedFixMap(feedbackStat.NeedFixMap),
		Comments:    feedbackStat.Comments,
	})
}