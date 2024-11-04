package entity

import "gorm.io/datatypes"

type Permission struct {
	ID          string         `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id,omitempty"`
	Name        string         `gorm:"uniqueIndex:permissions_name_unique" json:"name"`
	Description *string        `gorm:"nullable" json:"description"`
	Roles       []Role         `gorm:"many2many:permission_role" json:"roles,omitempty"`
	Metadata    datatypes.JSON `gorm:"type:jsonb;-" json:"metadata,omitempty"`
}
