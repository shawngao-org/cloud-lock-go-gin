package main

import (
	"cloud-lock-go-gin/config"
	"cloud-lock-go-gin/database"
	"cloud-lock-go-gin/influxdb"
	"cloud-lock-go-gin/redis"
	"cloud-lock-go-gin/server"
)

func main() {
	config.GetConfig()
	influxdb.ConnectInfluxDb()
	database.ConnectDb()
	redis.RedisConnect()
	server.HandleServer()
}
