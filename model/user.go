package model

import (
	"net/mail"
	"time"

	"github.com/microcosm-cc/bluemonday"
)

var policy = bluemonday.UGCPolicy()

type User struct {
	Id           int        `json:"id" db:"id"`
	Name         string     `json:"name" db:"name" binding:"required"`
	Mail         string     `json:"mail" db:"mail" binding:"required"`
	PasswordHash string     `json:"password" db:"password_hash" binding:"required"`
	Salt         string     `json:"-" db:"salt"`
	UserGender   *int       `json:"user_gender" db:"user_gender"`
	PreferGender *int       `json:"prefer_gender" db:"prefer_gender"`
	Description  *string    `json:"description" db:"description"`
	Age          *int       `json:"age"`
	Looking      *string    `json:"looking" db:"looking"`
	ImagePath    *string    `json:"image_path" db:"image_path"`
	Education    *string    `json:"education" db:"education"`
	Hobbies      *string    `json:"hobbies" db:"hobbies"`
	Birthday     *time.Time `json:"birthday" db:"birthday"`
	Online       bool       `json:"online" db:"online"`
	Tags         *[]string  `json:"tags" db:"tags"`
}

func (u *User) IsValid() bool {
	_, err := mail.ParseAddress(u.Mail)
	return err == nil && len(u.Name) > 1 && len(u.PasswordHash) > 5
}

func (u *User) CalculateAge() {
	if u.Birthday == nil {
		return
	}
	currentTime := time.Now()
	age := currentTime.Year() - u.Birthday.Year()

	if currentTime.YearDay() < u.Birthday.YearDay() {
		age--
	}
	u.Age = &age
}

func (u *User) Sanitize() {
	if u.Description != nil {
		*u.Description = policy.Sanitize(*u.Description)
	}
	if u.Looking != nil {
		*u.Looking = policy.Sanitize(*u.Looking)
	}
	if u.ImagePath != nil {
		*u.ImagePath = policy.Sanitize(*u.ImagePath)
	}
	if u.Education != nil {
		*u.Education = policy.Sanitize(*u.Education)
	}
	if u.Hobbies != nil {
		*u.Hobbies = policy.Sanitize(*u.Hobbies)
	}
	if u.Tags != nil {
		for i := 0; i < len(*u.Tags); i++ {
			(*u.Tags)[i] = policy.Sanitize((*u.Tags)[i])
		}
	}
	u.Name = policy.Sanitize(u.Name)
	u.Mail = policy.Sanitize(u.Mail)

	u.PasswordHash = ""
	u.Salt = ""
}
