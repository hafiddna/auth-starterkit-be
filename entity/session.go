package entity

import "github.com/hafiddna/auth-starterkit-be/entity/global"

type Session struct {
	global.Model
	UserID       string `gorm:"type:uuid;index;null" json:"user_id"`
	User         *User  `gorm:"foreignKey:UserID" json:"user,omitempty"`
	IPAddress    string `gorm:"type:varchar(45)" json:"ip_address"`
	UserAgent    string `gorm:"type:text" json:"user_agent"`
	Payload      string `gorm:"type:text" json:"payload"`
	LastActivity int64  `gorm:"type:integer;index" json:"last_activity"`
}
