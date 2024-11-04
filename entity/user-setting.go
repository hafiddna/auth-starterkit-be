package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserSetting struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Metadata EmbedJSON          `json:"metadata,omitempty" bson:"metadata,omitempty"`
}
