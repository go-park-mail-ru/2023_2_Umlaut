package constants

import "errors"

var (
	ErrAlreadyExists = errors.New("unique constraint error")
	ErrInvalidUser   = errors.New("invalid password or email")
	ErrMutualLike    = errors.New("mutual like")
	ErrNoFiles       = errors.New("no photos")
	ErrNoData        = errors.New("no data")
	ErrBannedUser    = errors.New("this user is blocked")
)
