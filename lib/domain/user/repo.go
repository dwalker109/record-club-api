package user

import (
	"context"
	"fmt"
	"github.com/dwalker109/record-club-api/lib/svc"
	"github.com/zmb3/spotify"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	coll = svc.Ctr.GetDBClient().Database("rc").Collection("users")
	ctx  = context.Background()
)

func GetUserFromSpotifyID(id string) (*Entity, error) {
	var ent Entity
	err := coll.FindOne(ctx, bson.M{"spotify_user": id}).Decode(&ent)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
	}

	return &ent, err
}

func AddUserFromSpotify(user *spotify.PrivateUser) (*Entity, error) {
	ent := Entity{
		Name:        user.DisplayName,
		SpotifyUser: user.ID,
	}

	result, err := coll.InsertOne(ctx, &ent)
	if err == nil {
		oid, ok := result.InsertedID.(primitive.ObjectID)
		if ok {
			ent.UserID = oid
		} else {
			err = fmt.Errorf("got data of type %T but wanted primitive.ObjectID", result)
		}
	}

	return &ent, err
}
