package routers

import (
	"github.com/gin-gonic/gin"
	"microservice-in-micro/part1/user-web/handler"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/hello", handler.Login)
	return router
}
