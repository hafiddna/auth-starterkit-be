package database

import (
	"fmt"
	"github.com/hafiddna/auth-starterkit-be/config"
	"github.com/redis/go-redis/v9"
)

func ConnectToRedis(db int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Config.App.Redis.Host, config.Config.App.Redis.Port),
		Password: config.Config.App.Redis.Password,
		DB:       db,
	})

	return rdb
}
