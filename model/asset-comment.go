package model

import "database/sql"

type AssetComment struct {
	Model
	AssetID string `gorm:"type:uuid;index" json:"asset_id"`
	//Asset      Asset          `gorm:"foreignKey:AssetID" json:"asset"`
	UserID   string         `gorm:"type:uuid;index" json:"user_id"`
	ParentID sql.NullString `gorm:"type:uuid" json:"parent_id"`
	//User       User           `gorm:"foreignKey:UserID" json:"user"`
	//Parent     *AssetComment  `gorm:"foreignKey:ParentID" json:"parent"`
	//Children   []AssetComment `gorm:"foreignKey:ParentID" json:"children"`
	Comment    string `gorm:"type:text" json:"comment"`
	IsResolved bool   `gorm:"type:boolean;default:false" json:"is_resolved"`
}

func (a *AssetComment) TableName() string {
	return "asset_comment"
}
