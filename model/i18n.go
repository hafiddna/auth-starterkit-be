package model

type I18n struct {
	ID string `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	//Translations []Translation `gorm:"foreignKey:I18nID" json:"translations"`
}

func (i *I18n) TableName() string {
	return "i18ns"
}
