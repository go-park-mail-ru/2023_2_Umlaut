package model

import "time"

type Admin struct {
	Id           int    `json:"id" db:"id"`
	Mail         string `json:"mail" db:"mail" binding:"required"`
	PasswordHash string `json:"password" db:"password_hash" binding:"required"`
	Salt         string `json:"-" db:"salt"`
}

type Recommendation struct {
	Id        int        `json:"id" db:"id"`
	UserId    int        `json:"user_id" db:"user_id"`
	Recommend *int       `json:"recommend" db:"recommend"`
	Show      bool       `json:"show" db:"show"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
}

func (u *Admin) Sanitize() {
	u.PasswordHash = ""
	u.Salt = ""
}


