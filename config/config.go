package config

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	Mongodb_uri_env = "MONGODB_URI"
	MONGODB_URI     = "mongodb://localhost:27017"
)

func ConnectToMongoDB() (*mongo.Client, error) {

	uri := os.Getenv(Mongodb_uri_env)
	if uri == "" {
		os.Setenv(Mongodb_uri_env, MONGODB_URI)
	}
	uri = os.Getenv(Mongodb_uri_env)
	if uri == "" {
		return nil, fmt.Errorf("MONGODB_URI is not set")
	}

	// clientOptions contains options like pool size,authentication
	clientOptions := options.Client().ApplyURI(MONGODB_URI)

	//Package context tells about like deadlines, cancellation signals.Further TODO() is used when it’s unclear which Context to use or it is not available
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	return client, nil

}
