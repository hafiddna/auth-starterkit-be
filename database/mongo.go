package database

import (
	"context"
	"fmt"
	"github.com/hafiddna/auth-starterkit-be/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongo() (database *mongo.Database, err error) {
	ctx := context.Background()

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s",
		config.Config.App.MongoDB.Username,
		config.Config.App.MongoDB.Password,
		config.Config.App.MongoDB.Host,
		config.Config.App.MongoDB.Port,
	)
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		return nil, err
	}

	database = client.Database(config.Config.App.MongoDB.Database)

	return database, nil
}
