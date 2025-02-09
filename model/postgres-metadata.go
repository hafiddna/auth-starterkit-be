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

func (m *Model) Updated(db *gorm.DB, userID string) {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	db.UpdateColumn("metadata", datatypes.JSONSet("metadata").Set("updated_by", userID).Set("updated_at", timestamp))
}

func (m *Model) SoftDelete(userID string) {
	// TODO: Implement SoftDelete method
	//timestamp := time.Now().UnixNano() / int64(time.Millisecond)
}
