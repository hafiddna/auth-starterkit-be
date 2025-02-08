package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserSetting struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID string             `json:"user_id" bson:"user_id"`
	//User     *User              `json:"user" bson:"-"`
	Metadata EmbedJSON `json:"metadata" bson:"metadata"`
}
