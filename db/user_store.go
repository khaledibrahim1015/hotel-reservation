package db

import (
	"context"

	"github.com/khaledibrahim1015/hotel-reservation/config"
	"github.com/khaledibrahim1015/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Concreste interface
type UserStore interface {
	GetUserByID(context.Context, string) (*types.User, error)
}

type MongoUserStore struct {
	mongoClient *mongo.Client
	coll        *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		mongoClient: client,
		coll:        client.Database(config.DBNAME).Collection(config.USERCOLL),
	}

}

func (store *MongoUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {

	var user types.User

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}
	if err := store.coll.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
