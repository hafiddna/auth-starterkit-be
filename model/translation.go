package model

type Translation struct {
	Model
	I18nID string `gorm:"type:uuid;not null" json:"i18n_id"`
	//I18n           I18n     `gorm:"foreignKey:I18nID" json:"i18n"`
	LanguageLocale string `gorm:"type:varchar(36);not null;index" json:"language_locale"`
	//Language       Language `gorm:"foreignKey:LanguageLocale;references:Locale" json:"language"`
	Value string `gorm:"type:text;not null" json:"value"`
}

func (t *Translation) TableName() string {
	return "translations"
}
