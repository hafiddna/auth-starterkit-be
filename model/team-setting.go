package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type TeamSetting struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	TeamID   string             `json:"team_id" bson:"team_id"`
	Team     *Team              `json:"team" bson:"-"`
	Metadata EmbedJSON          `json:"metadata" bson:"metadata"`
}
