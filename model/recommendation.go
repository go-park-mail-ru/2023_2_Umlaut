package model

import "time"

type Recommendation struct {
	Id        int        `json:"id" db:"id"`
	UserId    int        `json:"user_id" db:"user_id"`
	Recommend *int       `json:"recommend" db:"recommend"`
	Show      bool       `json:"show" db:"show"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
}

type RecommendationStatistic struct {
	AvgRecommend   float32 `json:"avg-recommend"`
	NPS            float32 `json:"nps"`
	RecommendCount []int32 `json:"recommend-count"`
}
