package global

import (
	"time"
)

type EmbedJSON struct {
	CreatedAt int64 `json:"created_at" bson:"created_at"`
	UpdatedAt int64 `json:"updated_at" bson:"updated_at"`
	DeletedAt int64 `json:"deleted_at" bson:"deleted_at"`

	CreatedBy string `json:"created_by" bson:"created_by"`
	UpdatedBy string `json:"updated_by" bson:"updated_by"`
	DeletedBy string `json:"deleted_by" bson:"deleted_by"`
}

func (d *EmbedJSON) SetCreated(createdBy string) {
	d.CreatedBy = createdBy
	d.CreatedAt = time.Now().UnixMilli()
}

func (d *EmbedJSON) SetUpdated(updatedBy string) {
	d.UpdatedBy = updatedBy
	d.UpdatedAt = time.Now().UnixMilli()
}

func (d *EmbedJSON) SetDeleted(deletedBy string) {
	d.DeletedBy = deletedBy
	d.DeletedAt = time.Now().UnixMilli()
}
