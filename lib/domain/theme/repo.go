package theme

import (
	"context"
	"github.com/dwalker109/record-club-api/lib/svc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = svc.Ctr.DBClient().Database("rc").Collection("themes")
var ctx = context.Background()

func GetAll() (*[]Entity, error) {
	var ents []Entity

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return &ents, err
	}

	err = cursor.All(ctx, &ents)

	return &ents, err
}

func GetOne(oid primitive.ObjectID) (*Entity, error) {
	var ent Entity
	err := collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&ent)

	return &ent, err
}

func AddOne(ent *Entity) error {
	_, err := collection.InsertOne(ctx, ent)

	return err
}
