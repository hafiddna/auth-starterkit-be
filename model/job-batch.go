package model

import "database/sql"

type JobBatch struct {
	Model
	Name         string         `gorm:"type:varchar(255);not null" json:"name"`
	TotalJobs    int            `gorm:"type:integer;not null" json:"total_jobs"`
	PendingJobs  int            `gorm:"type:integer;not null" json:"pending_jobs"`
	FailedJobs   int            `gorm:"type:integer;not null" json:"failed_jobs"`
	FailedJobIDs string         `gorm:"type:text;not null" json:"failed_job_ids"`
	Options      sql.NullString `gorm:"type:text" json:"options"`
	CancelledAt  sql.NullInt64  `gorm:"type:integer" json:"cancelled_at"`
	FinishedAt   sql.NullInt64  `gorm:"type:integer" json:"finished_at"`
}

func (j *JobBatch) TableName() string {
	return "job_batches"
}
