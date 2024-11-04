package entity

import "gorm.io/datatypes"

type Role struct {
	ID          string         `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id,omitempty"`
	Name        string         `gorm:"uniqueIndex:roles_name_unique" json:"name"`
	Description *string        `gorm:"nullable" json:"description"`
	Users       []User         `gorm:"many2many:role_user;-" json:"users,omitempty"`
	Permissions []Permission   `gorm:"many2many:permission_role" json:"permissions,omitempty"`
	Metadata    datatypes.JSON `gorm:"type:jsonb;-" json:"metadata,omitempty"`
}
