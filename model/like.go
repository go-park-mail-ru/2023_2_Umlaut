package model

type Like struct {
	LikedByUserId int  `json:"liked_by_user_id" db:"liked_by_user_id" swaggerignore:"true"`
	LikedToUserId int  `json:"liked_to_user_id" db:"liked_to_user_id"`
	IsLike        bool `json:"is_like" db:"is_like"`
}
