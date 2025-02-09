package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Model struct {
	ID       string         `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Metadata datatypes.JSON `gorm:"type:jsonb" json:"metadata"`
}

func (m *Model) Created(userID string) {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	m.Metadata = []byte(`{
		"created_by": "` + userID + `",
		"created_at": ` + strconv.FormatInt(timestamp, 10) + `,
		"updated_by": "` + userID + `",
		"updated_at": ` + strconv.FormatInt(timestamp, 10) + `,
		"deleted_by": null,
		"deleted_at": null
	}`)
}

func (m *Model) Updated(db *gorm.DB, userID string) *gorm.DB {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	return db.UpdateColumn("metadata", datatypes.JSONSet("metadata").Set("updated_by", userID).Set("updated_at", timestamp))
}

func (m *Model) SoftDelete(db *gorm.DB, userID string) *gorm.DB {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	return db.UpdateColumn("metadata", datatypes.JSONSet("metadata").Set("deleted_by", userID).Set("deleted_at", timestamp))
}

func OnlyTrashed(db *gorm.DB) *gorm.DB {
	return db.Where("metadata->>'deleted_at' IS NOT NULL").Where("metadata->>'deleted_by' IS NOT NULL")
}

func WithoutTrashed(db *gorm.DB) *gorm.DB {
	return db.Where("metadata->>'deleted_at' IS NULL").Where("metadata->>'deleted_by' IS NULL")
}
