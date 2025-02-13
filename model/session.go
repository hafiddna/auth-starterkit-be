package model

import "database/sql"

type Session struct {
	Model
	UserID sql.NullString `gorm:"type:uuid;index;null" json:"user_id"`
	//User         User           `gorm:"foreignKey:UserID" json:"user,omitempty"`
	IPAddress     sql.NullString `gorm:"type:varchar(45)" json:"ip_address"`
	UserAgent     sql.NullString `gorm:"type:text" json:"user_agent"`
	Payload       string         `gorm:"type:text" json:"payload"`
	LastActivity  int64          `gorm:"type:integer;index" json:"last_activity"`
	AppID         string         `gorm:"type:varchar(255);uniqueIndex" json:"app_id"`
	RememberToken sql.NullString `gorm:"type:varchar(100)" json:"remember_token"`
}

func (s *Session) TableName() string {
	return "sessions"
}
