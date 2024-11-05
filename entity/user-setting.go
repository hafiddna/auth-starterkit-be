package entity

import (
	"github.com/hafiddna/auth-starterkit-be/entity/global"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserSetting struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID   string             `json:"user_id" bson:"user_id"`
	User     *User              `json:"user" bson:"-"`
	Metadata global.EmbedJSON   `json:"metadata" bson:"metadata"`
}
