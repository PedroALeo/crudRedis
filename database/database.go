package database

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var (
	DB  *redis.Client
	CTX context.Context
)

// ConnectRedis start a connection with the redis database
func ConnectRedis() {
	CTX = context.Background()

	DB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
