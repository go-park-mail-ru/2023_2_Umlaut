package model

import "time"

type Like struct {
	Id            int       `json:"id" db:"id"`
	LikedByUserId int       `json:"liked_by_user_id" db:"liked_by_user_id"`
	LikedToUserId int       `json:"liked_to_user_id" db:"liked_to_user_id"`
	CommitedAt    time.Time `json:"commited_at" db:"commited_at"`
}
