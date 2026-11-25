package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()

	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Redis gagal connect:", err)
	}

	log.Println("Redis connected!")
}
