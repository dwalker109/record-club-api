package svc

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/oauth2"
)

type spotifyTokenMemCache map[primitive.ObjectID]*oauth2.Token

func (t *spotifyTokenMemCache) Store(id primitive.ObjectID, token *oauth2.Token) {
	(*t)[id] = token
}

func (t *spotifyTokenMemCache) Clear(id primitive.ObjectID) {
	delete(*t, id)
}
