package entity

import (
	"github.com/hafiddna/auth-starterkit-be/entity/global"
	"gorm.io/datatypes"
)

type Asset struct {
	global.Model
	OwnerID      string          `gorm:"type:uuid" json:"owner_id"`
	OwnerType    string          `gorm:"type:varchar(255)" json:"owner_type"`
	Owner        interface{}     `json:"owner,omitempty"`
	FolderID     string          `gorm:"type:uuid;null" json:"folder_id"`
	Folder       *Folder         `gorm:"foreignKey:FolderID" json:"folder,omitempty"`
	Name         string          `gorm:"type:varchar(255)" json:"name"`
	Path         string          `gorm:"type:varchar(255)" json:"path"`
	Size         int64           `gorm:"type:bigint;default:0" json:"size"`
	Type         string          `gorm:"type:varchar(255);default:'file';comment:'document, spreadsheet, presentation, form, image, pdf, video, shortcut, site, audio, drawing, archive, file'" json:"type"`
	Access       string          `gorm:"type:varchar(255);default:'viewer';comment:'viewer, editor, commenter'" json:"access"`
	BucketName   string          `gorm:"type:varchar(255)" json:"bucket_name"`
	IsPublic     bool            `gorm:"type:boolean;default:false" json:"is_public"`
	FileMetadata datatypes.JSON  `gorm:"type:jsonb;null" json:"file_metadata"`
	UserProfiles []*UserProfile  `gorm:"foreignKey:AvatarID" json:"user_profiles,omitempty"`
	Languages    []*Language     `gorm:"foreignKey:IconID" json:"languages,omitempty"`
	Shares       []*AssetShare   `gorm:"many2many:asset_share" json:"shares,omitempty"`
	Comments     []*AssetComment `gorm:"many2many:asset_comment" json:"comments,omitempty"`
	Tags         []*Tag          `gorm:"many2many:asset_tag" json:"tags,omitempty"`
}
