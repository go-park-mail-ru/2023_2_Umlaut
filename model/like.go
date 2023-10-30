package model

import "time"

type Like struct {
	LikedByUserId int       `json:"liked_by_user_id" db:"liked_by_user_id"`
	LikedToUserId int       `json:"liked_to_user_id" db:"liked_to_user_id"`
	CommittedAt   time.Time `json:"committed_at" db:"committed_at"`
}
