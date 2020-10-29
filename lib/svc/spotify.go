package svc

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/oauth2"
)

type SpotifyTokensMemCache map[primitive.ObjectID]*oauth2.Token

func (t *SpotifyTokensMemCache) Store(id primitive.ObjectID, token *oauth2.Token) {
	(*t)[id] = token
}

func (t *SpotifyTokensMemCache) Clear(id primitive.ObjectID) {
	delete(*t, id)
}
