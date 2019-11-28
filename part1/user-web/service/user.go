package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	//返回结果
	c.JSON(http.StatusOK, gin.H{
		"data": "hello-----world",
	})
}
