package entity

import "github.com/hafiddna/auth-starterkit-be/entity/global"

type Role struct {
	global.Model
	TeamID          string            `gorm:"type:uuid;index;nullable" json:"team_id"`
	Team            *Team             `gorm:"foreignKey:TeamID" json:"team,omitempty"`
	Name            string            `gorm:"uniqueIndex:roles_name_unique" json:"name"`
	Description     string            `gorm:"nullable" json:"description"`
	TeamInvitations []*TeamInvitation `gorm:"foreignKey:RoleID" json:"team_invitations,omitempty"`
	Users           []*User           `gorm:"many2many:user_role" json:"users,omitempty"`
	Permissions     []*Permission     `gorm:"many2many:role_permission" json:"permissions,omitempty"`
}
