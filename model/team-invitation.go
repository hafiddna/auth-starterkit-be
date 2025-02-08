package model

type TeamInvitation struct {
	Model
	TeamID string `gorm:"type:uuid;not null" json:"team_id"`
	//Team      Team   `gorm:"foreignKey:TeamID" json:"team"`
	Contact   string `gorm:"type:varchar(255);not null" json:"contact"`
	Type      string `gorm:"type:varchar(255);not null" json:"type"`
	Token     string `gorm:"type:varchar(255);not null;unique" json:"token"`
	ExpiresAt int64  `gorm:"type:int;not null" json:"expires_at"`
	RoleID    string `gorm:"type:uuid;not null" json:"role_id"`
	//Role      Role   `gorm:"foreignKey:RoleID" json:"role"`
}

func (t *TeamInvitation) TableName() string {
	return "team_invitations"
}
