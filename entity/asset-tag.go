package entity

import "github.com/hafiddna/auth-starterkit-be/entity/global"

type AssetTag struct {
	global.Model
	AssetID string `gorm:"type:uuid;index" json:"asset_id"`
	Asset   *Asset `gorm:"foreignKey:AssetID" json:"asset"`
	TagID   string `gorm:"type:uuid;index" json:"tag_id"`
	Tag     *Tag   `gorm:"foreignKey:TagID" json:"tag"`
}
