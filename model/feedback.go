package model

import "time"

type Feedback struct {
	Id         int        `json:"id" db:"id"`
	UserId     int        `json:"user_id" db:"user_id"`
	Rating     *int       `json:"rating" db:"rating"`
	Liked      *string    `json:"liked" db:"liked"`
	NeedFix    *string    `json:"need_fix" db:"need_fix"`
	CommentFix *string    `json:"comment_fix" db:"comment_fix"`
	Comment    *string    `json:"comment" db:"comment"`
	Show       bool       `json:"show" db:"show"`
	CreatedAt  *time.Time `json:"created_at" db:"created_at"`
}

type FeedbackStatistic struct {
	AvgRating   float32
	RatingCount []int32
	LikedMap    map[string]int32
	NeedFixMap  map[string]NeedFixObject
	Comments    []string
}

type NeedFixObject struct {
	Count      int32
	CommentFix []string
}
