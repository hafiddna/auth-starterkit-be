package entity

import "github.com/hafiddna/auth-starterkit-be/entity/global"

type UserRole struct {
	global.Model
	UserID string `gorm:"type:uuid;not null;index" json:"user_id"`
	User   User   `gorm:"foreignKey:UserID" json:"user"`
	RoleID string `gorm:"type:uuid;not null;index" json:"role_id"`
	Role   Role   `gorm:"foreignKey:RoleID" json:"role"`
}

func (u *UserRole) TableName() string {
	return "user_role"
}
