package database

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

type redisClient struct {
}

var c = context.Background()

// RedisClientInterface is an interface
type RedisClientInterface interface {
	RedisClient() *redis.Client
}

// NewRedisClient implements redisClient
func NewRedisClient() *redisClient {
	return &redisClient{}
}

func (r redisClient) RedisClient() *redis.Client {
	redisDataBase := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := redisDataBase.Ping(c).Result()

	if err != nil {
		log.Print("error")
	}

	return redisDataBase
}
