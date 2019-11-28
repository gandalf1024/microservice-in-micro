package basic

import (
	"microservice-in-micro/part1/basic/config"
	"microservice-in-micro/part1/basic/db"
	"microservice-in-micro/part1/basic/redis"
)

func Init() {
	config.Init()
	db.Init()
	redis.Init()
}
