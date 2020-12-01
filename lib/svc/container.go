package svc

import (
	"context"
	"github.com/go-chi/jwtauth"
	"go.mongodb.org/mongo-driver/mongo"
)

type container struct {
	mongoClient *mongo.Client
	jwtAuth     *jwtauth.JWTAuth
	authSvc     *authService
	acTokenSvc  *accessTokenService
	rfTokens    *refreshTokenMemCache
	spTokens    *spotifyTokenMemCache
}

func (c *container) DBClient() *mongo.Client {
	if c.mongoClient == nil {
		//c.mongoClient = mongoAtlasClient
		c.mongoClient = mongoLocalClient
	}
	return c.mongoClient
}

func (c *container) JWTAuth() *jwtauth.JWTAuth {
	if c.jwtAuth == nil {
		c.jwtAuth = JWTAuth
	}
	return c.jwtAuth
}

func (c *container) AuthService() *authService {
	if c.authSvc == nil {
		c.authSvc = &authService{}
	}
	return c.authSvc
}

func (c *container) AccessTokenService() *accessTokenService {
	if c.acTokenSvc == nil {
		c.acTokenSvc = &accessTokenService{}
	}
	return c.acTokenSvc
}

func (c *container) refreshTokenStore() *refreshTokenMemCache {
	if c.rfTokens == nil {
		cache := make(refreshTokenMemCache)
		c.rfTokens = &cache
	}
	return c.rfTokens
}

func (c *container) SpotifyTokenStore() *spotifyTokenMemCache {
	if c.spTokens == nil {
		cache := make(spotifyTokenMemCache)
		c.spTokens = &cache
	}
	return c.spTokens
}

func (c *container) Shutdown() {
	err := c.DBClient().Disconnect(context.Background())
	if err != nil {
		panic("Could not safely shutdown service container")
	}
}

// MakeContainer creates a new service container, suitable for use in tests
//func MakeContainer() *container {
//	return &container{}
//}

// Ctr is a service container suitable for general use
var Ctr = &container{}
