package model

type User struct {
	Id          int    `json:"-" db:"id"`
	Name        string `json:"name" binding:"required"`
	Username    string `json:"username" binding:"required"`
	Mail        string `json:"mail" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Description string `json:"description"`
	Age         int    `json:"age"`
	Looking     string `json:"looking"`
	Education   string `json:"education"`
	Hobbies     string `json:"hobbies"`
	Tags        string `json:"tags"`
}
