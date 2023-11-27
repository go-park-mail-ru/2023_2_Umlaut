package model

import "time"

type Complaint struct {
	Id             int        `json:"id" db:"id"`
	ReporterUserId int        `json:"reporter_user_id" db:"reporter_user_id"`
	ReportedUserId int        `json:"reported_user_id" db:"reported_user_id"`
	ComplaintType  string     `json:"complaint_type" db:"complaint_type"`
	ReportStatus   int        `json:"-" db:"report_status" swaggerignore:"true"`
	CreatedAt      *time.Time `json:"created_at" db:"created_at"`
}
