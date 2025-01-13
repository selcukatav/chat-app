package database

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

//var ctx = context.Background()

func Redis() {
	_ = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	fmt.Println("Connected to Redis!")

}
