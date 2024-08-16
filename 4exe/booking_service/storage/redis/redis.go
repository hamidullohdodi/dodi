package redis

import (
	"github.com/go-redis/redis/v8"
)

func connRD() *redis.Client {

	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // parolsiz ulanish
		DB:       0,  // default DB
	})
}
