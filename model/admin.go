package model

import "time"

type Admin struct {
	Id           int    `json:"id" db:"id"`
	Mail         string `json:"mail" db:"mail" binding:"required"`
	PasswordHash string `json:"password" db:"password_hash" binding:"required"`
	Salt         string `json:"-" db:"salt"`
}

type Statistic struct {
	Id         int        `json:"id" db:"id"`
	UserId     int        `json:"user_id" db:"user_id"`
	Rating     *int       `json:"rating" db:"rating"`
	Liked      *string    `json:"liked" db:"liked"`
	NeedFix    *string    `json:"need_fix" db:"need_fix"`
	CommentFix *string    `json:"comment_fix" db:"comment_fix"`
	Comment    *string    `json:"comment" db:"comment"`
	CreatedAt  *time.Time `json:"created_at" db:"created_at"`
}

type Recommendation struct {
	Id        int        `json:"id" db:"id"`
	UserId    int        `json:"user_id" db:"user_id"`
	Recommend *int       `json:"recommend" db:"recommend"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
}

func (u *Admin) Sanitize() {
	u.PasswordHash = ""
	u.Salt = ""
}
