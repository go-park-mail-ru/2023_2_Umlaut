package model

import "time"

type Complaint struct {
	Id              int `json:"id" db:"id"`
	ReporterId      int `json:"reporter_id" db:"reporter_id"`
	ReportedId      int `json:"reported_id" db:"reported_id"`
	ComplaintTypeId int `json:"complaint_type_id" db:"complaint_type_id"`
	ReportStatus    int `json:"report_status" db:"report_status"`
	TimeStamp       *time.Time `json:"timestamp" db:"timestamp"`
}
