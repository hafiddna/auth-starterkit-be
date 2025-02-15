package model

import "database/sql"

type Role struct {
	Model
	TeamID      sql.NullString `gorm:"type:uuid;index;nullable" json:"team_id"`
	Team        Team           `gorm:"foreignKey:TeamID" json:"team"`
	Name        string         `gorm:"uniqueIndex:roles_name_unique" json:"name"`
	DisplayName string         `gorm:"type:varchar(255)" json:"display_name"`
	Description sql.NullString `gorm:"nullable" json:"description"`
	//TeamInvitations []TeamInvitation `gorm:"foreignKey:RoleID" json:"team_invitations,omitempty"`
	Users       []User       `gorm:"many2many:user_role;joinForeignKey:RoleID;joinReferences:UserID" json:"users,omitempty"`
	Permissions []Permission `gorm:"many2many:role_permission" json:"permissions,omitempty"`
}

func (r *Role) TableName() string {
	return "roles"
}
