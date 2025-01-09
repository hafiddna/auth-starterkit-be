package global

import (
	"gorm.io/datatypes"
)

type Model struct {
	ID       string         `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Metadata datatypes.JSON `gorm:"type:jsonb" json:"metadata"`
}

func (m *Model) Created(userID string) {
	// set metadata created_by, created_at, updated_by, updated_at
}

func (m *Model) Updated(userID string) {
	// set metadata updated_by, updated_at
}

func (m *Model) SoftDelete(userID string) {
	// set metadata deleted_by, deleted_at
}
