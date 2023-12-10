package model

import "time"

type Recommendation struct {
	Id        int        `json:"id" db:"id"`
	UserId    int        `json:"user_id" db:"user_id"`
	Rating    *int       `json:"rating" db:"rating"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
}

type RecommendationStatistic struct {
	AvgRecommend   float32 `json:"avg-recommend"`
	NPS            float32 `json:"nps"`
	RecommendCount []int32 `json:"recommend-count"`
}
