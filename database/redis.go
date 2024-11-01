package database

import (
	"context"
	"fmt"
	"github.com/hafiddna/auth-starterkit-be/config"
	"github.com/redis/go-redis/v9"
	"time"
)

type Redis interface {
	Connect() *redis.Client
	Set(key string, value interface{}, expiration int) error
	Get(key string) (string, error)
	Delete(key string) error
}

type redisStruct struct {
	config config.CfgStruct
	ctx    context.Context
	db     int
}

func NewRedis(config config.CfgStruct, ctx context.Context, db int) Redis {
	return &redisStruct{
		config: config,
		ctx:    ctx,
		db:     db,
	}
}

func (r *redisStruct) Connect() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", r.config.App.Redis.Host, r.config.App.Redis.Port),
		Password: r.config.App.Redis.Password,
		DB:       r.db,
	})

	return rdb
}

func (r *redisStruct) Set(key string, value interface{}, expiration int) error {
	rdb := r.Connect()

	err := rdb.Set(r.ctx, key, value, time.Duration(expiration)*time.Second).Err()

	if err != nil {
		return err
	}

	return nil
}

func (r *redisStruct) Get(key string) (string, error) {
	rdb := r.Connect()

	val, err := rdb.Get(r.ctx, key).Result()

	if err != nil {
		return "", err
	}

	return val, nil
}

func (r *redisStruct) Delete(key string) error {
	rdb := r.Connect()

	err := rdb.Del(r.ctx, key).Err()

	if err != nil {
		return err
	}

	return nil
}
