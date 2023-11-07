package model

import "errors"

var AlreadyExists = errors.New("unique constraint error")

var InvalidUser = errors.New("invalid password or email")

var MutualLike = errors.New("mutual like")
