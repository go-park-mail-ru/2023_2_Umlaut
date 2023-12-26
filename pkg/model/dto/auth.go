package dto

type SignInInput struct {
	Mail     string `json:"mail" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignUpInput struct {
	Name      string  `json:"name" binding:"required"`
	Mail      string  `json:"mail" binding:"required"`
	Password  string  `json:"password" binding:"required"`
	InvitedBy *string `json:"invited_by"`
}