package repositories

import (
	"context"
	"os"
	"quasarproject/database"
	"strconv"
	"time"
)

var (
	ctx context.Context = context.Background()
)

// RepositoryInterface is an interface
type RepositoryInterface interface {
	Set(key string, value string) error
	Get(key string) (string, error)
	Ping() error
}

type repository struct {
	client database.ClientInterface
}

// NewRepository implements RepositoryInterface
func NewRepository(client database.ClientInterface) RepositoryInterface {
	return &repository{
		client: client,
	}
}

func (r repository) Ping() error {
	redisClient := r.client.Client().RedisClient()
	defer redisClient.Close()
	_, err := redisClient.Do(ctx, "PING").Result()
	return r.checkError(err)
}

func (r repository) checkError(err error) error {
	if err != nil {
		return err
	}
	return nil
}

func (r repository) Set(key string, value string) error {
	redisClient := r.client.Client().RedisClient()
	defer redisClient.Close()

	seconds, _ := strconv.Atoi(os.Getenv("EXPIRATION_TIME"))
	err := redisClient.Set(ctx, key, value, time.Duration(seconds)*time.Second).Err()

	return r.checkError(err)
}

func (r repository) Get(key string) (string, error) {
	redisClient := r.client.Client().RedisClient()
	defer redisClient.Close()

	value, err := redisClient.Get(ctx, key).Result()

	if err != nil {
		return "", err
	}

	return value, nil
}
