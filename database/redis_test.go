package database

import (
	"context"
	"github.com/hafiddna/auth-starterkit-be/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	redisCtx = context.Background()

	redisConfig = config.NewConfig().GetConfig()

	redisDB = NewRedis(redisConfig, redisCtx, 15)
)

func TestRedisStruct_Connect(t *testing.T) {
	client := redisDB.Connect()

	assert.NotNil(t, client)
}

func TestRedisStruct_GetSetDelete(t *testing.T) {
	key := "test"
	value := "test"

	err := redisDB.Set(key, value, 10)
	assert.Nil(t, err)

	val, err := redisDB.Get(key)
	assert.Nil(t, err)
	assert.Equal(t, value, val)

	err = redisDB.Delete(key)
	assert.Nil(t, err)
}
