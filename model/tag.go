package model

type Tag struct {
	Model
	Name string `gorm:"type:varchar(255);unique" json:"name"`
	//Assets []Asset `gorm:"many2many:asset_tag" json:"assets"`
}

func (t *Tag) TableName() string {
	return "tags"
}
