package entity

import (
	"github.com/hafiddna/auth-starterkit-be/entity/global"
	"gorm.io/datatypes"
)

type Asset struct {
	global.Model
	OwnerID      string         `gorm:"type:uuid" json:"owner_id"`
	Owner        *User          `gorm:"foreignKey:OwnerID" json:"owner"`
	FolderID     string         `gorm:"type:uuid" json:"folder_id"`
	Folder       *Folder        `gorm:"foreignKey:FolderID" json:"folder"`
	Name         string         `gorm:"type:varchar(255)" json:"name"`
	Path         string         `gorm:"type:varchar(255)" json:"path"`
	Size         int64          `gorm:"type:bigint;default:0" json:"size"`
	Type         string         `gorm:"type:varchar(255);default:'file';comment:'document, spreadsheet, presentation, form, image, pdf, video, shortcut, site, audio, drawing, archive, file'" json:"type"`
	Access       string         `gorm:"type:varchar(255);default:'viewer';comment:'viewer, editor, owner'" json:"access"`
	BucketName   string         `gorm:"type:varchar(255)" json:"bucket_name"`
	IsPublic     bool           `gorm:"type:boolean;default:false" json:"is_public"`
	FileMetadata datatypes.JSON `gorm:"type:jsonb;null" json:"file_metadata"`
}
