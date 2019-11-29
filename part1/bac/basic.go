package bac

import (
	"bac/config"
	"bac/db"
	"bac/redis"
)

func Init() {
	config.Init()
	db.Init()
	redis.Init()
}
