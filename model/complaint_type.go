package model

type ComplaintType struct {
	Id   int    `json:"id" db:"id"`
	TypeName string `json:"type_name" db:"type_name"`
}
