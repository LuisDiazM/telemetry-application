package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/cache/v9"
	"github.com/phuslu/log"
	"github.com/redis/go-redis/v9"
)

type RedisImp struct {
	Settings *RedisSettings
	Cache    *cache.Cache
	Ctx      *context.Context
}

func NewRedisImp(settings *RedisSettings) *RedisImp {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": settings.RedisHost,
		},
		Password: settings.RedisPass,
		PoolSize: 20,
	})

	mycache := cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})
	ctx := context.TODO()
	return &RedisImp{Settings: settings, Cache: mycache, Ctx: &ctx}
}

func (cacheMgt *RedisImp) Set(prefix string, key string, value interface{}, ttl time.Duration) {
	data, _ := json.Marshal(value)
	err := cacheMgt.Cache.Set(&cache.Item{
		Ctx:   *cacheMgt.Ctx,
		Key:   prefix + ":" + key,
		Value: data,
		TTL:   ttl,
	})
	if err != nil {
		log.Error().Msg(err.Error())
	}
}

func (cacheMgt *RedisImp) Get(prefix string, key string, data interface{}) (err error) {
	cachedHolder := make([]byte, 5)
	err = cacheMgt.Cache.Get(*cacheMgt.Ctx, prefix+":"+key, &cachedHolder)
	if err == nil {
		_ = json.Unmarshal(cachedHolder, &data)
	}
	return
}

func (cacheMgt *RedisImp) Delete(prefix string, key string) (err error) {
	err = cacheMgt.Cache.Delete(*cacheMgt.Ctx, prefix+":"+key)
	return
}
