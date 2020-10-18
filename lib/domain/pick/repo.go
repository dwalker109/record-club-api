package pick

import (
	"context"
	"github.com/dwalker109/record-club-api/lib/svc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection = svc.Ctr.GetDBClient().Database("rc").Collection("picks")
var ctx = context.Background()

func GetThemePicks(themeID primitive.ObjectID) (*[]Entity, error) {
	return doGet(&bson.D{{"theme_id", themeID}})
}

func GetThemePicksForOwner(themeID, ownerID primitive.ObjectID) (*[]Entity, error) {
	return doGet(&bson.D{{"theme_id", themeID}, {"owner_id", ownerID}})
}

func doGet(filter *bson.D) (*[]Entity, error) {
	var ents []Entity

	cursor, err := collection.Find(ctx, *filter)
	if err != nil {
		return &ents, err
	}

	err = cursor.All(ctx, &ents)

	return &ents, err
}

func AddOrUpdateThemePicksForOwner(ent *Entity) error {
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"theme_id", ent.ThemeID}, {"owner_id", ent.OwnerID}}
	_, err := collection.UpdateOne(ctx, filter, bson.D{{"$set", ent}}, opts)

	return err
}
