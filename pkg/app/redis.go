package app

import (
	"5z7Game/config"
	"5z7Game/pkg/utils"
	"github.com/go-redis/redis"
	"sync"
)

var client = new(RedisClient)

type RedisClient struct {
	once sync.Once
	instance *redis.Client
}

// connect
func (r *RedisClient) connect(options *redis.Options) error {
	client.instance = redis.NewClient(options)

	return client.instance.Ping().Err()
}

// Redis return redis client
func Redis() *redis.Client {
	client.once.Do(func() {
		utils.SecurePanic(client.connect(config.Redis().Options()))
	})

	return client.instance
}

