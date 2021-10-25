package component

import (
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

var Cache *cache.Cache

func init() {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": "127.0.0.1:6379",
		},
	})

	Cache = cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})
}
