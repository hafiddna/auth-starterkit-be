package entity

import (
	"database/sql"
	"github.com/hafiddna/auth-starterkit-be/entity/global"
)

type Language struct {
	global.Model
	Name         string         `gorm:"type:varchar(255);not null" json:"name"`
	Code         string         `gorm:"type:varchar(255);not null" json:"code"`
	Locale       string         `gorm:"type:varchar(255);not null;unique;index" json:"locale"`
	IsRTL        bool           `gorm:"type:boolean;default:false" json:"is_rtl"`
	IconID       sql.NullString `gorm:"type:uuid;null" json:"icon_id"`
	Icon         Asset          `gorm:"foreignKey:IconID" json:"icon,omitempty"`
	Translations []Translation  `gorm:"foreignKey:LanguageLocale;references:Locale" json:"translations,omitempty"`
}

func (l *Language) TableName() string {
	return "languages"
}
