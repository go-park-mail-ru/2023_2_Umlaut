package model

import (
	"net/mail"
)

type User struct {
	Id           int        `json:"id" db:"id"`
	Name         string     `json:"name" db:"name" binding:"required"`
	Mail         string     `json:"mail" db:"mail" binding:"required"`
	PasswordHash string     `json:"password,omitempty" db:"password_hash" binding:"required"`
	Salt         string     `json:"-" db:"salt"`
	UserGender   NullString `json:"user_gender" db:"user_gender"`
	PreferGender NullString `json:"prefer_gender" db:"prefer_gender"`
	Description  NullString `json:"description" db:"description"`
	Age          NullInt64  `json:"age" db:"age"`
	Looking      NullString `json:"looking" db:"looking"`
	Education    NullString `json:"education" db:"education"`
	Hobbies      NullString `json:"hobbies" db:"hobbies"`
	Tags         NullString `json:"tags" db:"tags"`
}

func (u *User) Sanitize() {
	u.PasswordHash = ""
}

func (u *User) IsValid() bool {
	_, err := mail.ParseAddress(u.Mail)
	return err == nil && len(u.Name) > 1 && len(u.PasswordHash) > 5
}
