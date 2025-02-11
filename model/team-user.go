package model

type TeamUser struct {
	Model
	TeamID   string `gorm:"type:uuid;not null;index" json:"team_id"`
	Team     Team   `gorm:"foreignKey:TeamID" json:"team"`
	UserID   string `gorm:"type:uuid;not null;index" json:"user_id"`
	User     User   `gorm:"foreignKey:UserID" json:"user"`
	IsActive bool   `gorm:"type:boolean;default:false" json:"is_active"`
}

func (t *TeamUser) TableName() string {
	return "team_user"
}
