package model

import "time"

type Complaint struct {
	Id              int        `json:"id" db:"id"`
	ReporterUserId  int        `json:"reporter_user_id" db:"reporter_user_id"`
	ReportedUserId  int        `json:"reported_user_id" db:"reported_user_id"`
	ComplaintTypeId int        `json:"complaint_type_id" db:"complaint_type_id"`
	CreatedAt       *time.Time `json:"created_at" db:"created_at"`
}
