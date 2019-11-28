package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	us "microservice-in-micro/part1/user-web/proto/user"
	"net/http"
)

var (
	serviceClient us.UserService
)

// Error 错误结构体
type Error struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

func Init() {
	serviceClient = us.NewUserService("mu.micro.book.srv.user", client.DefaultClient)
}

func Login(c *gin.Context) {
	uaername := c.Request.FormValue("userName")

	// 调用后台服务
	rsp, err := serviceClient.QueryUserByName(context.TODO(), &us.Request{
		UserName: uaername,
	})
	if err != nil {
		http.Error(c.Writer, err.Error(), 500)
		return
	}

	u := rsp.GetUser()
	//返回结果
	c.JSON(http.StatusOK, gin.H{
		"data":     u.CreatedTime,
		"username": u.Name,
		"password": u.Pwd,
	})
}
