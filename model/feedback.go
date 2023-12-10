package model

import "time"

type Feedback struct {
	Id        int        `json:"id" db:"id"`
	UserId    int        `json:"user_id" db:"user_id"`
	Rating    *int       `json:"rating" db:"rating"`
	Liked     *string    `json:"liked" db:"liked"`
	NeedFix   *string    `json:"need_fix" db:"need_fix"`
	Comment   *string    `json:"comment" db:"comment"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
}

type FeedbackStatistic struct {
	AvgRating   float32                  `json:"avg-rating"`
	RatingCount []int32                  `json:"rating-count"`
	LikedMap    map[string]int32         `json:"liked-map"`
	NeedFixMap  map[string]NeedFixObject `json:"need-fix-map"`
	Comments    []string                 `json:"comments"`
}

type NeedFixObject struct {
	Count      int32    `json:"count"`
	CommentFix []string `json:"comment_fix"`
}
