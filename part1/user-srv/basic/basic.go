package basic

import (
	"microservice-in-micro/part1/user-srv/basic/config"
	"microservice-in-micro/part1/user-srv/basic/db"
)

func Init() {
	config.Init()
	db.Init()
}
