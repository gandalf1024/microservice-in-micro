package routers

import (
	"github.com/gin-gonic/gin"
	"microservice-in-micro/part1/user-web/service"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/hello", service.Login)
	return router
}
