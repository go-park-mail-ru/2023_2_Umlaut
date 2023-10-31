package model

type Dialog struct {
	Id      int `json:"id" db:"id"`
	User1Id int `json:"user1_id" db:"user1_id"`
	User2Id int `json:"user2_id" db:"user2_id"`
}
