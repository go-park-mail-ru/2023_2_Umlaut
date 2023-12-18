package dto

import "github.com/go-park-mail-ru/2023_2_Umlaut/pkg/model/core"

type FilterParams struct {
	UserId int      `json:"user_id"`
	MinAge int      `json:"min_age"`
	MaxAge int      `json:"max_age"`
	Tags   []string `json:"tags"`
}

type FeedData struct {
	User        core.User `json:"user"`
	LikeCounter int  `json:"like_counter"`
}
