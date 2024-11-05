package entity

import "github.com/hafiddna/auth-starterkit-be/entity/global"

type TeamUser struct {
	global.Model
	TeamID string `gorm:"type:uuid;not null" json:"team_id"`
	Team   *Team  `gorm:"foreignKey:TeamID" json:"team"`
	UserID string `gorm:"type:uuid;not null" json:"user_id"`
	User   *User  `gorm:"foreignKey:UserID" json:"user"`
}
