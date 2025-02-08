package model

type PasswordResetToken struct {
	Model
	UserID string `gorm:"type:uuid;not null;index" json:"user_id"`
	//User      User   `gorm:"foreignKey:UserID" json:"user"`
	Contact   string `gorm:"type:varchar(255);not null" json:"contact"`
	Type      string `gorm:"type:varchar(255);not null" json:"type"`
	Token     string `gorm:"type:varchar(255);not null" json:"token"`
	Attempts  int    `gorm:"type:integer;not null;default:0" json:"attempts"`
	ExpiresAt int64  `gorm:"type:integer;not null" json:"expires_at"`
}

func (p *PasswordResetToken) TableName() string {
	return "password_reset_tokens"
}
