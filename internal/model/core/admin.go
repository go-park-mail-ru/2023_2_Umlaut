package core

type Admin struct {
	Id           int    `json:"id" db:"id"`
	Mail         string `json:"mail" db:"mail" binding:"required"`
	PasswordHash string `json:"password" db:"password_hash" binding:"required"`
	Salt         string `json:"-" db:"salt"`
}

func (u *Admin) Sanitize() {
	u.PasswordHash = ""
	u.Salt = ""
}
