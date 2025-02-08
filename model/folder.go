package model

import "database/sql"

type Folder struct {
	Model
	//Assets    []Asset        `gorm:"foreignKey:FolderID" json:"assets"`
	OwnerType sql.NullString `gorm:"type:varchar(255)" json:"owner_type"`
	OwnerID   sql.NullString `gorm:"type:uuid" json:"owner_id"`
	Owner     interface{}    `json:"owner,omitempty"`
	ParentID  sql.NullString `gorm:"type:uuid" json:"parent_id"`
	//Parent    *Folder        `gorm:"foreignKey:ParentID" json:"parent"`
	//Children  []Folder       `gorm:"foreignKey:ParentID" json:"children"`
	Name string `gorm:"type:varchar(255)" json:"name"`
}

func (f *Folder) TableName() string {
	return "folders"
}
