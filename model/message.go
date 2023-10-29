package model

import "time"

type Message struct {
	Id          int        `json:"id" db:"id"`
	SenderId    int        `json:"sender_id" db:"sender_id"`
	DialogId    int        `json:"dialog_id" db:"dialog_id"`
	MessageText string     `json:"message_text" db:"message_text"`
	TimeStamp   *time.Time `json:"timestamp" db:"timestamp"`
}
