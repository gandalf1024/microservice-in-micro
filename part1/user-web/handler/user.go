package handler

import (
	au "auth/proto/auth"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"
	"net/http"
	us "user_srv/proto/user"
)

var (
	serviceClient us.UserService
	authClient    au.Service
)

// Error 错误结构体
type Error struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

func Init() {
	serviceClient = us.NewUserService("mu.micro.book.srv.user", client.DefaultClient)
	authClient = au.NewService("mu.micro.book.srv.auth", client.DefaultClient)
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

	// 生成token
	rsp2, err := authClient.MakeAccessToken(context.TODO(), &au.Request{
		UserId:   uint64(rsp.User.Id),
		UserName: rsp.User.Name,
	})

	if err != nil {
		log.Logf("[Login] 创建token失败，err：%s", err)
		http.Error(c.Writer, err.Error(), 500)
		return
	}

	log.Logf("[Login] token %s", rsp2.Token)
	// 同时将token写到cookies中
	c.Header("set-cookie", "application/json; charset=utf-8")
	// 过期30分钟
	c.SetCookie("remember-me-token", rsp2.Token, 90000, "/", "", true, true)

	u := rsp.GetUser()
	//返回结果
	c.JSON(http.StatusOK, gin.H{
		"data":     u.CreatedTime,
		"username": u.Name,
		"password": u.Pwd,
	})
}

func Logout(c *gin.Context) {

}
