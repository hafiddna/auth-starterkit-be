package entity

import (
	"time"
)

type EmbedJSON struct {
	CreatedAt string `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DeletedAt string `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`

	CreatedBy string `json:"created_by,omitempty" bson:"created_by,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty" bson:"updated_by,omitempty"`
	DeletedBy string `json:"deleted_by,omitempty" bson:"deleted_by,omitempty"`
}

func (d *EmbedJSON) SetCreated(createdBy string) {
	d.CreatedBy = createdBy
	d.CreatedAt = time.Now().Format(time.RFC3339)
}

func (d *EmbedJSON) SetUpdated(updatedBy string) {
	d.UpdatedBy = updatedBy
	d.UpdatedAt = time.Now().Format(time.RFC3339)
}

func (d *EmbedJSON) SetDeleted(deletedBy string) {
	d.DeletedBy = deletedBy
	d.DeletedAt = time.Now().Format(time.RFC3339)
}
