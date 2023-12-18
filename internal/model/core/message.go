package core

import "time"

type Message struct {
	Id          *int       `json:"id" db:"id"`
	SenderId    *int       `json:"sender_id" db:"sender_id"`
	RecipientId *int       `json:"recipient_id" db:"recipient_id"`
	DialogId    *int       `json:"dialog_id" db:"dialog_id"`
	Text        *string    `json:"message_text" db:"message_text"`
	IsRead      *bool      `json:"is_read" db:"is_read"`
	CreatedAt   *time.Time `json:"created_at" db:"created_at"`
}
