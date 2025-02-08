package model

type FailedJob struct {
	Model
	UUID       string `gorm:"type:uuid;unique" json:"uuid"`
	Connection string `gorm:"type:text" json:"connection"`
	Queue      string `gorm:"type:text" json:"queue"`
	Payload    string `gorm:"type:text" json:"payload"`
	Exception  string `gorm:"type:text" json:"exception"`
	FailedAt   string `gorm:"type:timestamp" json:"failed_at"`
}

func (f *FailedJob) TableName() string {
	return "failed_jobs"
}
