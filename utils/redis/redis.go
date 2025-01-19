package redisDB

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

var (
	Client *redis.Client
	Ctx    = context.Background()
)

// Initialize initializes the Redis client
func Initialize() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Change the address if required
		Password: "",
		DB:       0,
	})

	_, err := Client.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
}

// Close closes the Redis connection
func Close() {
	if err := Client.Close(); err != nil {
		log.Printf("Error closing Redis connection: %v", err)
	}
}
