package database

import (
	"context"
	"github.com/hafiddna/auth-starterkit-be/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	mongoCtx = context.Background()

	mongoConfig = config.NewConfig().GetConfig()

	mongoTest = NewMongoDB(mongoConfig, mongoCtx)
)

func TestMongoDB_ConnectDisconnect(t *testing.T) {
	client, err := mongoTest.Connect()

	assert.Nil(t, err)

	err = mongoTest.Disconnect(client)

	assert.Nil(t, err)
}
