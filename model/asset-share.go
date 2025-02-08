package model

type AssetShare struct {
	Model
	AssetID string `gorm:"type:uuid;index" json:"asset_id"`
	//Asset   Asset  `gorm:"foreignKey:AssetID" json:"asset"`
	UserID string `gorm:"type:uuid;index" json:"user_id"`
	//User    User   `gorm:"foreignKey:UserID" json:"user"`
	Access string `gorm:"type:varchar(255);default:'viewer';comment:'viewer, editor, commenter'" json:"access"`
}

func (a *AssetShare) TableName() string {
	return "asset_share"
}
