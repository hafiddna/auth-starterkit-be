package model

type TeamUser struct {
	Model
	TeamID string `gorm:"type:uuid;not null" json:"team_id"`
	//Team   *Team  `gorm:"foreignKey:TeamID" json:"team"`
	UserID string `gorm:"type:uuid;not null" json:"user_id"`
	//User   *User  `gorm:"foreignKey:UserID" json:"user"`
}

func (t *TeamUser) TableName() string {
	return "team_user"
}
