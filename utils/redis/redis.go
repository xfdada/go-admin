package redis

import (
	"fmt"
	"github.com/go-redis/redis"
)

var Redis *redis.Client

func init() {
	NewClient()
}

func NewClient() *redis.Client {
	if Redis != nil {
		return Redis
	}

	Redis = redis.NewClient(&redis.Options{
		Addr:     "182.92.234.23:6379",
		Password: "123456", // no password set
		DB:       0,        // use default DB
	})

	_, err := Redis.Ping().Result()
	if err != nil {
		fmt.Printf("redis connection failed: %v\n", err.Error())
	}

	return Redis
}
