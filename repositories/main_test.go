package repositories

import (
	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v8"
)

func mockRedis() *miniredis.Miniredis {
	s, err := miniredis.Run()

	if err != nil {
		panic(err)
	}

	return s
}

func newTestRedis() *redis.Client {

	s := mockRedis()

	client := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})

	return client
}
