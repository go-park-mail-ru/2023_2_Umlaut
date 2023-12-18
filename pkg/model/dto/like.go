package dto

type PremiumLike struct {
	LikedByUserId int       `json:"liked_by_user_id"`
	ImagePaths    *[]string `json:"image_paths"`
}
