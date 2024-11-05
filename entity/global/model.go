package global

import (
	"gorm.io/datatypes"
)

type Model struct {
	ID       string         `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Metadata datatypes.JSON `gorm:"type:jsonb" json:"metadata"`
}

func (m *Model) Created(userID string) {
}

func (m *Model) Updated(userID string) {
}

func (m *Model) SoftDelete(userID string) {
}
