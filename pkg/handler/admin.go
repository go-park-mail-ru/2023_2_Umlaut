package handler

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/go-park-mail-ru/2023_2_Umlaut/model"
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/admin/proto"
	"github.com/go-park-mail-ru/2023_2_Umlaut/utils"
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
// @ID Recommendation
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
// @ID FeedFeedback
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
// @ID CSAT
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
// @ID RecommendationStatistic
// @Accept  json
// @Produce  json
// @Success 200 {object} ClientResponseDto[model.RecommendationStatistic]
// @Failure 500 {object} ClientResponseDto[string]
// @Router /api/v1/admin/recommendation [post]
func (h *Handler) getRecommendationStatistic(w http.ResponseWriter, r *http.Request) {
	recommend, err := h.adminMicroservice.GetRecommendationStatistic(r.Context(), &proto.AdminEmpty{})
	if err != nil {
		statusCode, message := parseError(err)
		newErrorClientResponseDto(r.Context(), w, statusCode, message)
		return
	}

	NewSuccessClientResponseDto(r.Context(), w, &model.RecommendationStatistic{
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
// @Router /api/v1/admin/feedback [post]
func (h *Handler) getFeedbackStatistic(w http.ResponseWriter, r *http.Request) {
	feedbackStat, err := h.adminMicroservice.GetFeedbackStatistic(r.Context(), &proto.AdminEmpty{})
	if err != nil {
		statusCode, message := parseError(err)
		newErrorClientResponseDto(r.Context(), w, statusCode, message)
		return
	}

	NewSuccessClientResponseDto(r.Context(), w, &model.FeedbackStatistic{
		AvgRating:   feedbackStat.AvgRating,
		RatingCount: feedbackStat.RatingCount,
		LikedMap:    getLikedMap(feedbackStat.LikedMap),
		NeedFixMap:  getNeedFixMap(feedbackStat.NeedFixMap),
		Comments:    feedbackStat.Comments,
	})
}

func getLikedMap(likedMap []*proto.LikedMap) map[string]int32 {
	result := make(map[string]int32)
	for _, item := range likedMap {
		result[item.Liked] = item.Count
	}
	return result
}

func getNeedFixMap(needFixMap []*proto.NeedFixMap) map[string]model.NeedFixObject {
	result := make(map[string]model.NeedFixObject)
	for _, item := range needFixMap {
		result[item.NeedFix] = model.NeedFixObject{
			Count:      item.NeedFixObject.Count,
			CommentFix: item.NeedFixObject.CommentFix,
		}
	}
	return result
}
