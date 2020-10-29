package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entity struct {
	UserID      primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	SpotifyUser string             `bson:"spotify_user,omitempty"`
}
