package model

type Dialog struct {
	Id      int `json:"id" db:"id"`
	User1Id int `json:"user_1_id" db:"user_1_id"`
	User2Id int `json:"user_2_id" db:"user_2_id"`
}
