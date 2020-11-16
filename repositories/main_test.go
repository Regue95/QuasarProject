package repositories

import (
	"log"
	"quasarproject/database"
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v8"
)

type clientMock struct {
}

type miniRedisClientMock struct {
}

func (m miniRedisClientMock) RedisClient() *redis.Client {
	miniRedis, err := miniredis.Run()

	if err != nil {
		log.Print("error")
	}

	miniRedisClient := redis.NewClient(&redis.Options{
		Addr: miniRedis.Addr(),
	})

	return miniRedisClient
}

func NewClientMock() database.ClientInterface {
	return &clientMock{}
}

func (c clientMock) Client() database.RedisClientInterface {
	redis := miniRedisClientMock{}
	return redis
}

func TestRedisPing(t *testing.T) {
	dataBaseClient := NewClientMock()
	repo := NewRepository(dataBaseClient)
	result := repo.Ping()

	if result != nil {
		log.Fatal("Error TestRedisPing")
	}
}
func TestRedisGet(t *testing.T) {
	dataBaseClient := NewClientMock()
	repo := NewRepository(dataBaseClient)
	result, _ := repo.Get("satelite")
	expected := ""

	if result != expected {
		log.Fatal(result, " es distinto a: ", expected)
	}
}

func TestRedisSet(t *testing.T) {
	dataBaseClient := NewClientMock()
	repo := NewRepository(dataBaseClient)
	result := repo.Set("satelite", "satelite")

	if result != nil {
		log.Fatal("Error TestRedisSet")
	}
}
