package entity

type Session struct {
	ID           string  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id,omitempty"`
	UserID       string  `gorm:"type:uuid;index" json:"user_id"`
	User         User    `gorm:"foreignKey:UserID" json:"user"`
	IPAddress    string  `gorm:"type:varchar(45)" json:"ip_address"`
	UserAgent    *string `gorm:"type:text" json:"user_agent"`
	Payload      string  `gorm:"type:text" json:"payload"`
	LastActivity int64   `gorm:"type:integer" json:"last_activity"`
}
