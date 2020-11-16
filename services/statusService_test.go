package services

import (
	"log"
	"quasarproject/database"
	"quasarproject/repositories"
	"quasarproject/responses"
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

func TestGetPing(t *testing.T) {
	dataBaseClient := NewClientMock()

	repo := repositories.NewRepository(dataBaseClient)

	statService := NewStatusService(repo)

	result, _ := statService.GetPing()

	expected := responses.Status{Status: "OK"}

	if result != expected {
		log.Fatal(result, " es distinto a: ", expected)
	}
}

func TestGetPingDB(t *testing.T) {
	dataBaseClient := NewClientMock()

	repo := repositories.NewRepository(dataBaseClient)

	statService := NewStatusService(repo)

	result := statService.PingDB()

	if result != nil {
		log.Fatal("Error TestGetPingDB")
	}
}
