package model

type I18n struct {
	Model
	//Translations []Translation `gorm:"foreignKey:I18nID" json:"translations"`
}

func (i *I18n) TableName() string {
	return "i18ns"
}
