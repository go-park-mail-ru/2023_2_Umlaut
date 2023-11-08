package model

import "time"

type Like struct {
	LikedByUserId int       `json:"liked_by_user_id" db:"liked_by_user_id" swaggerignore:"true"`
	LikedToUserId int       `json:"liked_to_user_id" db:"liked_to_user_id"`
	CommittedAt   time.Time `json:"-" db:"committed_at"`
}
