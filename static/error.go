package static

import "errors"

var (
	ErrAlreadyExists = errors.New("unique constraint error")
	ErrInvalidUser = errors.New("invalid password or email")
	ErrMutualLike = errors.New("mutual like")
	ErrNoFiles = errors.New("no photos")
)