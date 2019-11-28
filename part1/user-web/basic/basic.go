package basic

import (
	"microservice-in-micro/part1/user-web/basic/config"
	"microservice-in-micro/part1/user-web/basic/db"
)

func Init() {
	config.Init()
	db.Init()
}
