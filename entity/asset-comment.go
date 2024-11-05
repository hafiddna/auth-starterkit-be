package entity

import "github.com/hafiddna/auth-starterkit-be/entity/global"

type AssetComment struct {
	global.Model
	AssetID    string          `gorm:"type:uuid;index" json:"asset_id"`
	Asset      *Asset          `gorm:"foreignKey:AssetID" json:"asset"`
	UserID     string          `gorm:"type:uuid;index" json:"user_id"`
	User       *User           `gorm:"foreignKey:UserID" json:"user"`
	ParentID   string          `gorm:"type:uuid" json:"parent_id"`
	Parent     *AssetComment   `gorm:"foreignKey:ParentID" json:"parent"`
	Children   []*AssetComment `gorm:"foreignKey:ParentID" json:"children"`
	Comment    string          `gorm:"type:text" json:"comment"`
	IsResolved bool            `gorm:"type:boolean;default:false" json:"is_resolved"`
}
