package db

import (
	"context"
	"fmt"

	"github.com/khaledibrahim1015/hotel-reservation/config"
	"github.com/khaledibrahim1015/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Concreste interface
type UserStore interface {
	GetUserByID(context.Context, string) (*types.User, error)
	GetUsers(context.Context) ([]*types.User, error)
	InsertUser(context.Context, *types.User) (*types.User, error)
	UpdateUser(context.Context, string, *types.UpdateUserParam) error
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

func (store *MongoUserStore) GetUsers(ctx context.Context) ([]*types.User, error) {
	var users []*types.User

	cursor, err := store.coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var user types.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)

	}
	return users, nil

}

func (store *MongoUserStore) InsertUser(ctx context.Context, user *types.User) (*types.User, error) {

	result, err := store.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	// Ensure the InsertedID is of type primitive.ObjectID before converting
	if objectID, ok := result.InsertedID.(primitive.ObjectID); ok {
		user.ID = objectID.Hex()
	} else {
		return nil, fmt.Errorf("failed to convert InsertedID to ObjectID")
	}

	return user, nil
}

func (store *MongoUserStore) UpdateUser(ctx context.Context, id string, user *types.UpdateUserParam) error {

	// First, check if the user exists
	_, err := store.GetUserByID(ctx, id)
	if err != nil {
		return fmt.Errorf("user not found with ID %s: %v", id, err)
	}

	// Create an update document with only the non-nil fields
	update := bson.M{}
	if user.FirstName != "" {
		update["firstName"] = user.FirstName
	}
	if user.LastName != "" {
		update["lastName"] = user.LastName
	}
	if user.Email != "" {
		update["email"] = user.Email
	}

	// If no fields to update, return early
	if len(update) == 0 {
		return fmt.Errorf("no update fields provided")
	}

	// perform update
	// Convert string ID to ObjectID
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid user ID format: %v", err)
	}
	filter := bson.M{"_id": objId}
	updateResult, err := store.coll.UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}

	// This check is somewhat redundant due to the earlier existence check,
	// but it's good to keep for potential race conditions
	if updateResult.ModifiedCount == 0 {
		return fmt.Errorf("no user found with ID %s", id)
	}
	return nil
}
