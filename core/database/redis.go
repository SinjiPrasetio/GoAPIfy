package database

import (
	"GoAPIfy/core/helper"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

// InitRedis initializes a Redis client and returns it, or returns nil if Redis is not enabled.
func InitRedis() *redis.Client {
	// Get the ENABLE_REDIS environment variable and convert it to a boolean value.
	enableRedisStr := os.Getenv("ENABLE_REDIS")
	enableRedis, err := strconv.ParseBool(enableRedisStr)
	if err != nil {
		log.Fatal("Error converting ENABLE_REDIS to boolean.")
	}

	// If Redis is not enabled, return nil.
	if !enableRedis {
		return nil
	}

	// Get the Redis host, port, and password from the environment variables, or use default values.
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	if redisHost == "" {
		redisHost = "localhost"
	}

	if redisPort == "" {
		redisPort = "6379"
	}

	// Create a new Redis client with the specified options.
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPassword,
		DB:       0,
	})

	// Ping the Redis server to check if the connection is working.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}

	// Print a message to indicate that the Redis client has been initialized.
	fmt.Println(helper.ColorizeCmd(helper.Green, "Connected to Redis..."))

	// Return the Redis client.
	return client
}
