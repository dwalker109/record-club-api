package svc

import (
	"context"
	"github.com/go-chi/jwtauth"
	"go.mongodb.org/mongo-driver/mongo"
)

type container struct {
	mongoDBClient *mongo.Client
	tokenAuth     *jwtauth.JWTAuth
}

func (c *container) GetDBClient() *mongo.Client {
	if c.mongoDBClient == nil {
		c.mongoDBClient = MongoClient
	}
	return c.mongoDBClient
}

func (c *container) GetJWTTokenAuth() *jwtauth.JWTAuth {
	if c.tokenAuth == nil {
		c.tokenAuth = JWTTokenAuth
	}
	return c.tokenAuth
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
