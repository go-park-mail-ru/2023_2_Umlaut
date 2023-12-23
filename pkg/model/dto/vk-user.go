package dto

type VkUser struct {
	Id        int    `json:"id"`
	Name      string `json:"first_name"`
	Email     string `json:"email"`
	Photo     string `json:"photo_max"`
	Sex       int    `json:"sex"`
	BirthDate string `json:"bdate"`
	InvitedBy string `json:"invited_by"`
}
