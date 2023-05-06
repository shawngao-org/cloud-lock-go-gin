package redis

import (
	"cloud-lock-go-gin/config"
	"github.com/go-redis/redis/v8"
)

var RedisClient = redisConnect()

func redisConnect() *redis.Client {
	host := config.Conf.Database.Redis.Host
	port := config.Conf.Database.Redis.Port
	password := config.Conf.Database.Redis.Password
	db := config.Conf.Database.Redis.Db
	return redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       db,
	})
}
