package entity

import "gorm.io/datatypes"

type Asset struct {
	ID string `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id,omitempty"`
	//OwnerID *string `gorm:"type:uuid" json:"owner_id"`
	//OwnerType    *string        `gorm:"type:varchar(255)" json:"owner_type"`
	OwnerID      string         `gorm:"type:uuid;index" json:"owner_id_id"`
	Owner        User           `gorm:"foreignKey:OwnerID" json:"owner"`
	Name         string         `gorm:"type:varchar(255)" json:"name"`
	Type         string         `gorm:"type:enum('image', 'video', 'pdf', 'file');not null" json:"type"`
	Access       string         `gorm:"type:enum('public', 'private');default:'private'" json:"access"`
	BucketType   string         `gorm:"type:enum('public', 'private');default:'private'" json:"bucket_type"`
	Path         string         `gorm:"type:varchar(255)" json:"path"`
	Bytes        float64        `gorm:"type:decimal(20,0);default:0" json:"bytes"`
	FileMetadata datatypes.JSON `gorm:"type:jsonb;null" json:"file_metadata,omitempty"`
	Metadata     datatypes.JSON `gorm:"type:jsonb;-" json:"metadata,omitempty"`
}
