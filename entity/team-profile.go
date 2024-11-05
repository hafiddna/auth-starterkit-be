package entity

import (
	"github.com/hafiddna/auth-starterkit-be/entity/global"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TeamProfile struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	TeamID   string             `json:"team_id" bson:"team_id"`
	Team     *Team              `json:"team" bson:"-"`
	Metadata global.EmbedJSON   `json:"metadata" bson:"metadata"`
}
