package database

import (
	"github.com/go-redis/redis/v8"
)

var (
	DB *redis.Client
)

func ConnectRedis() {
	DB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
