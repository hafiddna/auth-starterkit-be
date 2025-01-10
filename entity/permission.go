package entity

import (
	"database/sql"
	"github.com/hafiddna/auth-starterkit-be/entity/global"
)

type Permission struct {
	global.Model
	Name        string         `gorm:"type:varchar(255);uniqueIndex:permissions_name_unique" json:"name"`
	Description sql.NullString `gorm:"type:varchar(255);nullable" json:"description"`
	Roles       []Role         `gorm:"many2many:role_permission" json:"roles,omitempty"`
}

func (p *Permission) TableName() string {
	return "permissions"
}
