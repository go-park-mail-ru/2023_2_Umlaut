package utils

import (
	"github.com/go-park-mail-ru/2023_2_Umlaut/pkg/constants"
	"net/http"
	"strconv"
	"strings"

	feedProto "github.com/go-park-mail-ru/2023_2_Umlaut/pkg/microservices/feed/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ParseError(err error) (int, string) {
	code, ok := status.FromError(err)
	if ok {
		switch code.Code() {
		case codes.InvalidArgument:
			return http.StatusBadRequest, code.Message()
		case codes.Unauthenticated:
			return http.StatusUnauthorized, code.Message()
		case codes.Internal:
			return http.StatusInternalServerError, code.Message()
		case codes.NotFound:
			return http.StatusNotFound, code.Message()
		case codes.PermissionDenied:
			return http.StatusForbidden, code.Message()
		case codes.DataLoss:
			return http.StatusRequestURITooLong, code.Message()
		case codes.ResourceExhausted:
			return http.StatusPaymentRequired, code.Message()
		}
	}
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}

	return 200, ""
}

func ParseQueryParams(r *http.Request) *feedProto.FilterParams {
	minAge, _ := strconv.Atoi(r.URL.Query().Get("min_age"))
	maxAge, _ := strconv.Atoi(r.URL.Query().Get("max_age"))
	tags := strings.Split(r.URL.Query().Get("tags"), ",")
	return &feedProto.FilterParams{
		UserId: int32(r.Context().Value(constants.KeyUserID).(int)),
		MinAge: int32(minAge),
		MaxAge: int32(maxAge),
		Tags:   tags,
	}
}
