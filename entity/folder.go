package entity

import "github.com/hafiddna/auth-starterkit-be/entity/global"

type Folder struct {
	global.Model
	Assets   []*Asset  `gorm:"foreignKey:FolderID" json:"assets"`
	OwnerID  string    `gorm:"type:uuid" json:"owner_id"`
	Owner    *User     `gorm:"foreignKey:OwnerID" json:"owner"`
	ParentID string    `gorm:"type:uuid" json:"parent_id"`
	Parent   *Folder   `gorm:"foreignKey:ParentID" json:"parent"`
	Children []*Folder `gorm:"foreignKey:ParentID" json:"children"`
}
