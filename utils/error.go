package utils

import (
	"net/http"

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
		}
	}
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}

	return 200, ""
}
