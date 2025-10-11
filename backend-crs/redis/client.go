package redis

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func NewRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if err := rdb.Ping(Ctx).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis Database %s", err)
	}
	log.Println("Connected to the Redis Database Successfully")
	return rdb
}
