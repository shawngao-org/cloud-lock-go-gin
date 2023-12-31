package redis

import (
	"cloud-lock-go-gin/config"
	"github.com/go-redis/redis/v8"
	"sync"
)

var (
	redisMutex sync.Mutex
	Client     *redis.Client
)

func RedisConnect() {
	host := config.Conf.Database.Redis.Host
	port := config.Conf.Database.Redis.Port
	password := config.Conf.Database.Redis.Password
	db := config.Conf.Database.Redis.Db
	redisMutex.Lock()
	Client = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       db,
	})
	redisMutex.Unlock()
}
