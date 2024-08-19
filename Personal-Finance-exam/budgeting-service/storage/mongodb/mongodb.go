package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Connect(ctx context.Context) (*mongo.Database, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongodb:27017"))
	if err != nil {
		return nil, err
	}

	db := client.Database("exam")

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return db, nil

}