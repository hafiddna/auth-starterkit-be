package entity

import (
	"database/sql"
	"github.com/hafiddna/auth-starterkit-be/entity/global"
)

type Job struct {
	global.Model
	Queue       string        `gorm:"type:varchar(255);index" json:"queue"`
	Payload     string        `gorm:"type:text" json:"payload"`
	Attempts    int           `gorm:"type:integer" json:"attempts"`
	ReservedAt  sql.NullInt64 `gorm:"type:integer" json:"reserved_at"`
	AvailableAt int64         `gorm:"type:integer" json:"available_at"`
}

func (j *Job) TableName() string {
	return "jobs"
}
