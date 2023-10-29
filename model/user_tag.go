package model

type UserTag struct {
	Id     int `json:"id" db:"id"`
	UserId int `json:"user_id" db:"user_id"`
	TagId  int `json:"tag_id" db:"tag_id"`
}
