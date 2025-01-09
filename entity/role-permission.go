package entity

import "github.com/hafiddna/auth-starterkit-be/entity/global"

type RolePermission struct {
	global.Model
	PermissionID string      `gorm:"type:uuid;not null" json:"permission_id"`
	Permission   *Permission `gorm:"foreignKey:PermissionID" json:"permission"`
	RoleID       string      `gorm:"type:uuid;not null" json:"role_id"`
	Role         *Role       `gorm:"foreignKey:RoleID" json:"role"`
}

func (RolePermission) TableName() string {
	return "role_permission"
}
