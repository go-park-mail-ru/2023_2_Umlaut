package handler

import (
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/model/core"
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/model/dto"
	"github.com/go-park-mail-ru/2023_2_Umlaut/internal/utils"
	"net/http"
)

// @Summary get user for feed
// @Tags feed
// @ID feed
// @Accept  json
// @Produce  json
// @Param min_age query integer false "Minimum age filter"
// @Param max_age query integer false "Maximum age filter"
// @Param tags query string false "Tags filter"
// @Success 200 {object} ClientResponseDto[model.FeedData]
// @Failure 500 {object} ClientResponseDto[string]
// @Router /api/v1/feed [get]
func (h *Handler) feed(w http.ResponseWriter, r *http.Request) {
	feed, err := h.feedMicroservice.Feed(
		r.Context(),
		utils.ParseQueryParams(r),
	)

	if err != nil {
		statusCode, message := utils.ParseError(err)
		dto.NewErrorClientResponseDto(r.Context(), w, statusCode, message)
		return
	}
	birthday := feed.User.Birthday.AsTime()
	preferGender := int(feed.User.PreferGender)
	age := int(feed.User.Age)

	dto.NewSuccessClientResponseDto(r.Context(), w, dto.FeedData{
		User: core.User{
			Id:           int(feed.User.Id),
			Name:         feed.User.Name,
			PreferGender: &preferGender,
			Description:  &feed.User.Description,
			Age:          &age,
			Looking:      &feed.User.Looking,
			Education:    &feed.User.Education,
			Hobbies:      &feed.User.Hobbies,
			Birthday:     &birthday,
			Tags:         &feed.User.Tags,
			ImagePaths:   &feed.User.ImagePaths,
		},
		LikeCounter: int(feed.LikeCounter),
	})
}
