package database

import (
	"context"
	"fmt"
	"github.com/hafiddna/auth-starterkit-be/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB interface {
	Connect() (*mongo.Client, error)
	Disconnect(*mongo.Client) error
}

type mongoDB struct {
	config config.CfgStruct
	ctx    context.Context
}

func NewMongoDB(config config.CfgStruct, ctx context.Context) MongoDB {
	return &mongoDB{
		config: config,
		ctx:    ctx,
	}
}

func (m *mongoDB) Connect() (*mongo.Client, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s",
		m.config.App.MongoDB.Username,
		m.config.App.MongoDB.Password,
		m.config.App.MongoDB.Host,
		m.config.App.MongoDB.Port,
	)
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(m.ctx, clientOptions)

	if err != nil {
		return nil, err
	}

	err = client.Ping(m.ctx, nil)

	if err != nil {
		return nil, err
	}

	return client, nil
}

func (m *mongoDB) Disconnect(client *mongo.Client) error {
	err := client.Disconnect(m.ctx)

	if err != nil {
		return err
	}

	return nil
}
