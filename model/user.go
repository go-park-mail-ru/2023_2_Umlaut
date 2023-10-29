package model

import (
	"net/mail"
	"time"
)

type User struct {
	Id           int        `json:"id" db:"id"`
	Name         string     `json:"name" db:"name" binding:"required"`
	Mail         string     `json:"mail" db:"mail" binding:"required"`
	PasswordHash string     `json:"password,omitempty" db:"password_hash" binding:"required"`
	Salt         string     `json:"-" db:"salt"`
	UserGender   *string    `json:"user_gender" db:"user_gender"`
	PreferGender *string    `json:"prefer_gender" db:"prefer_gender"`
	Description  *string    `json:"description" db:"description"`
	Age          *int       `json:"age" db:"age"`
	Looking      *string    `json:"looking" db:"looking"`
	ImagePath    *string    `json:"image_path" db:"image_path"`
	Education    *string    `json:"education" db:"education"`
	Hobbies      *string    `json:"hobbies" db:"hobbies"`
	Tags         *string    `json:"tags" db:"tags"`
	Birthday     *time.Time `json:"birthday" db:"birthday"`
	Online       bool       `json:"online" db:"online"`
}

func (u *User) Sanitize() {
	u.PasswordHash = ""
}

func (u *User) IsValid() bool {
	_, err := mail.ParseAddress(u.Mail)
	return err == nil && len(u.Name) > 1 && len(u.PasswordHash) > 5
}
