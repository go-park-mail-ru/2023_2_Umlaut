package model

type FilterParams struct {
	UserId int      `json:"user_id"`
	MinAge int      `json:"min_age"`
	MaxAge int      `json:"max_age"`
	Tags   []string `json:"tags"`
}
