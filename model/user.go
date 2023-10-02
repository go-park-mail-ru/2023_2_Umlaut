package model

type User struct {
	Id           int     `json:"-" db:"id"`
	Name         string  `json:"name" db:"name" binding:"required"`
	Mail         string  `json:"mail" db:"mail" binding:"required"`
	PasswordHash string  `json:"password,omitempty" db:"password_hash" binding:"required"`
	Salt         string  `json:"-" db:"salt"`
	UserGender   *string `json:"user_gender" db:"user_gender"`
	PreferGender *string `json:"prefer_gender" db:"prefer_gender"`
	Description  *string `json:"description" db:"description"`
	Age          *int    `json:"age" db:"age"`
	Looking      *string `json:"looking" db:"looking"`
	Education    *string `json:"education" db:"education"`
	Hobbies      *string `json:"hobbies" db:"hobbies"`
	Tags         *string `json:"tags" db:"tags"`
}

func (u *User) Sanitize() {
	u.PasswordHash = ""
}
