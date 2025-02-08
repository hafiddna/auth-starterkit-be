package model

type AssetTag struct {
	Model
	AssetID string `gorm:"type:uuid;index" json:"asset_id"`
	//Asset   Asset  `gorm:"foreignKey:AssetID" json:"asset"`
	TagID string `gorm:"type:uuid;index" json:"tag_id"`
	//Tag     Tag    `gorm:"foreignKey:TagID" json:"tag"`
}

func (a *AssetTag) TableName() string {
	return "asset_tag"
}
