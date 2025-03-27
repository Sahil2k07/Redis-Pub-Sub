package configs

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func RedisInit() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       4,
		Protocol: 2,
	})

	err := checkConnection(client)
	if err != nil {
		panic("No Redis Connection")
	}

	RedisClient = client
}

func checkConnection(client *redis.Client) error {
	ctx := context.Background()

	_, err := client.Ping(ctx).Result()

	return err
}
