package svc

import (
	"context"
	"github.com/dwalker109/record-club-api/lib/svc/tokens"
	"github.com/go-chi/jwtauth"
	"go.mongodb.org/mongo-driver/mongo"
)

type container struct {
	mongoClient *mongo.Client
	jwtAuth     *jwtauth.JWTAuth
	rfTokens    *tokens.RefreshTokensMemCache
	spTokens    *SpotifyTokensMemCache
}

func (c *container) GetDBClient() *mongo.Client {
	if c.mongoClient == nil {
		//c.mongoClient = mongoAtlasClient
		c.mongoClient = mongoLocalClient
	}
	return c.mongoClient
}

func (c *container) GetJWTAuth() *jwtauth.JWTAuth {
	if c.jwtAuth == nil {
		c.jwtAuth = JWTAuth
	}
	return c.jwtAuth
}

func (c *container) GetRefreshTokensStore() *tokens.RefreshTokensMemCache {
	if c.rfTokens == nil {
		cache := make(tokens.RefreshTokensMemCache)
		c.rfTokens = &cache
	}
	return c.rfTokens
}

func (c *container) GetSpotifyTokensStore() *SpotifyTokensMemCache {
	if c.spTokens == nil {
		cache := make(SpotifyTokensMemCache)
		c.spTokens = &cache
	}
	return c.spTokens
}

func (c *container) Shutdown() {
	c.GetDBClient().Disconnect(context.Background())
}

// MakeContainer creates a new service container, suitable for use in tests
func MakeContainer() *container {
	return &container{}
}

// Ctr is a service container suitable for general use
var Ctr = &container{}
