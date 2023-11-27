package handler

import (
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/feed/proto"
	"github.com/go-park-mail-ru/2023_2_Umlaut/static"
	"net/http"
	"strconv"
	"strings"
)

// @Summary get user for feed
// @Tags feed
// @ID feed
// @Accept  json
// @Produce  json
// @Param min_age query integer false "Minimum age filter"
// @Param max_age query integer false "Maximum age filter"
// @Param tags query string false "Tags filter"
// @Success 200 {object} ClientResponseDto[model.User]
// @Failure 500 {object} ClientResponseDto[string]
// @Router /api/v1/feed [get]
func (h *Handler) feed(w http.ResponseWriter, r *http.Request) {
	user, err := h.feedMicroservice.Feed(
		r.Context(),
		parseQueryParams(r),
	)

	if err != nil {
		statusCode, message := parseError(err)
		newErrorClientResponseDto(r.Context(), w, statusCode, message)
		return
	}
	NewSuccessClientResponseDto(r.Context(), w, user)
}

func parseQueryParams(r *http.Request) *proto.FilterParams {
	minAge, _ := strconv.Atoi(r.URL.Query().Get("min_age"))
	maxAge, _ := strconv.Atoi(r.URL.Query().Get("max_age"))
	tags := strings.Split(r.URL.Query().Get("tags"), ",")
	return &proto.FilterParams{
		UserId: int32(r.Context().Value(static.KeyUserID).(int)),
		MinAge: int32(minAge),
		MaxAge: int32(maxAge),
		Tags:   tags,
	}
}
