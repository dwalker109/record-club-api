package svc

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type container struct {
	DBClient *mongo.Client
}

func (c *container) GetDBClient() *mongo.Client {
	if c.DBClient == nil {
		c.DBClient = MongoClient
	}
	return c.DBClient
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
