package model

import (
	"encoding/json"
	"github.com/hafiddna/auth-starterkit-be/helper"
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
	if userID == "" {
		userID = "system"
	}
	m.Metadata = []byte(`{
		"created_by": "` + userID + `",
		"created_at": ` + strconv.FormatInt(timestamp, 10) + `,
		"updated_by": "` + userID + `",
		"updated_at": ` + strconv.FormatInt(timestamp, 10) + `,
		"deleted_by": null,
		"deleted_at": null
	}`)
}

func (m *Model) Updated(userID string) {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	if userID == "" {
		userID = "system"
	}

	var metadata map[string]interface{}
	err := helper.JSONUnmarshal(m.Metadata, &metadata)
	if err != nil {
		metadata = map[string]interface{}{}
	}

	metadata["updated_by"] = userID
	metadata["updated_at"] = timestamp

	updatedMetadata, err := json.Marshal(metadata)
	if err != nil {
		updatedMetadata = m.Metadata
	}

	m.Metadata = updatedMetadata
}

func (m *Model) SoftDelete(userID string) {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	if userID == "" {
		userID = "system"
	}

	var metadata map[string]interface{}
	err := helper.JSONUnmarshal(m.Metadata, &metadata)
	if err != nil {
		metadata = map[string]interface{}{}
	}

	metadata["deleted_by"] = userID
	metadata["deleted_at"] = timestamp

	updatedMetadata, err := json.Marshal(metadata)
	if err != nil {
		updatedMetadata = m.Metadata
	}

	m.Metadata = updatedMetadata
}

func OnlyTrashed(db *gorm.DB) *gorm.DB {
	return db.Where("metadata->>'deleted_at' IS NOT NULL").Where("metadata->>'deleted_by' IS NOT NULL")
}

func WithoutTrashed(db *gorm.DB) *gorm.DB {
	return db.Where("metadata->>'deleted_at' IS NULL").Where("metadata->>'deleted_by' IS NULL")
}
