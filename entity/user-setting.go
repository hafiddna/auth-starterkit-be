package entity

import (
	"github.com/hafiddna/auth-starterkit-be/entity/global"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserSetting struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Metadata global.EmbedJSON   `json:"metadata,omitempty" bson:"metadata,omitempty"`
}
