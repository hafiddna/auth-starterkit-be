package model

import "database/sql"

type Permission struct {
	Model
	Name        string         `gorm:"type:varchar(255);uniqueIndex:permissions_name_unique" json:"name"`
	Description sql.NullString `gorm:"type:varchar(255);nullable" json:"description"`
	//Roles       []Role         `gorm:"many2many:role_permission" json:"roles,omitempty"`
}

func (p *Permission) TableName() string {
	return "permissions"
}
