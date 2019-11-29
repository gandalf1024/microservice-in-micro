package handler

import (
	au "auth/proto/auth"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"
	"microservice-in-micro/part1/user-web/gin/err_code"
	"microservice-in-micro/part1/user-web/gin/res"
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

func Login(ctx *gin.Context) {
	resp := res.Gin{C: ctx}
	uaername := ctx.Request.FormValue("userName")

	// 调用后台服务
	rsp, err := serviceClient.QueryUserByName(context.TODO(), &us.Request{
		UserName: uaername,
	})
	if err != nil {
		resp.Response(http.StatusInternalServerError, err_code.ERROR, err)
		return
	}

	// 生成token
	rsp2, err := authClient.MakeAccessToken(context.TODO(), &au.Request{
		UserId:   uint64(rsp.User.Id),
		UserName: rsp.User.Name,
	})

	if err != nil {
		log.Logf("[Login] 创建token失败，err：%s", err)
		resp.Response(http.StatusInternalServerError, err_code.ERROR, err)
		return
	}

	log.Logf("[Login] token %s", rsp2.Token)
	// 同时将token写到cookies中
	ctx.Header("set-cookie", "application/json; charset=utf-8")
	//ctx.Header("remember-me-token", rsp2.Token)
	// 过期30分钟
	ctx.SetCookie("remember-me-token", rsp2.Token, 90000, "/", "", false, true)

	u := rsp.GetUser()
	//返回结果
	resp.Response(http.StatusOK, err_code.SUCCESS, gin.H{
		"data":     u.CreatedTime,
		"username": u.Name,
		"password": u.Pwd,
	})
}

func Logout(ctx *gin.Context) {
	resp := res.Gin{C: ctx}
	tokenCookie, err := ctx.Cookie("remember-me-token")
	if err != nil {
		log.Logf("token获取失败")
		resp.Response(http.StatusInternalServerError, err_code.ERROR, err)
		return
	}

	// 删除token
	_, err = authClient.DelUserAccessToken(context.TODO(), &au.Request{
		Token: tokenCookie,
	})
	if err != nil {
		resp.Response(http.StatusInternalServerError, err_code.ERROR, err)
		return
	}

	resp.Response(http.StatusOK, err_code.SUCCESS, gin.H{
		"status": "success",
	})
}
