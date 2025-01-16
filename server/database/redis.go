package database

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func Redis() *redis.Client {
	if redisClient == nil {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
		fmt.Println("Connected to Redis!")
	}
	return redisClient
}