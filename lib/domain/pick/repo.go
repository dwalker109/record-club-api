package pick

import (
	"context"
	"github.com/dwalker109/record-club-api/lib/svc"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

var collection = svc.Ctr.GetDBClient().Database("rc").Collection("picks")
var ctx = context.Background()

func GetAll() (*[]Entity, error) {
	var ents []Entity

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return &ents, err
	}

	if err := cursor.All(ctx, &ents); err != nil {
		return &ents, err
	}

	return &ents, nil
}

func GetOne(id string) (*Entity, error) {
	p, _ := uuid.Parse(id)
	var ent Entity

	err := collection.FindOne(ctx, bson.M{"_id": p}).Decode(&ent)
	if err != nil {
		return &ent, err
	}

	return &ent, nil

}

func AddOne(e *Entity) error {
	_, err := collection.InsertOne(ctx, e)

	return err
}

func DeleteOne(id string) error {
	p, _ := uuid.Parse(id)
	_, err := collection.DeleteOne(ctx, bson.M{"_id": p})

	return err
}
