package entity

import "github.com/hafiddna/auth-starterkit-be/entity/global"

type Folder struct {
	global.Model
	Assets    []*Asset    `gorm:"foreignKey:FolderID" json:"assets"`
	OwnerID   string      `gorm:"type:uuid" json:"owner_id"`
	OwnerType string      `gorm:"type:varchar(255)" json:"owner_type"`
	Owner     interface{} `json:"owner,omitempty"`
	ParentID  string      `gorm:"type:uuid" json:"parent_id"`
	Parent    *Folder     `gorm:"foreignKey:ParentID" json:"parent"`
	Children  []*Folder   `gorm:"foreignKey:ParentID" json:"children"`
	Name      string      `gorm:"type:varchar(255)" json:"name"`
}
