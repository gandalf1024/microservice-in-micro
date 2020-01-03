package handler

import (
	au "auth/proto/auth"
	"context"
	hystrix_go "github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-plugins/wrapper/breaker/hystrix"
	"go.uber.org/zap"
	"microservice-in-micro/part1/user-web/gin/err_code"
	"microservice-in-micro/part1/user-web/gin/res"
	"net/http"
	z "plugins/zap"
	us "user_srv/proto/user"
)

var (
	log = z.GetLogger()
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
	hystrix_go.DefaultVolumeThreshold = 1
	hystrix_go.DefaultErrorPercentThreshold = 1
	cl := hystrix.NewClientWrapper()(client.DefaultClient) // 包装
	_ = cl.Init(
		client.Retries(3),
		//为了调试看log方便，始终返回true, nil，即会一直重试直至重试次数用尽
		//client.Retry(func(ctx context.Context, req client.Request, retryCount int, err error) (bool, error) {
		//	//log.Info(req.Method(), zap.Any("retryCount:", retryCount))
		//	fmt.Println("===========================>>>", req.Method(), retryCount)
		//	return true, nil
		//}),
	)
	serviceClient = us.NewUserService("mu.micro.book.srv.user", cl)
	authClient = au.NewService("mu.micro.book.srv.auth", cl)
}

func Login(ctx *gin.Context) {
	resp := res.Gin{C: ctx}
	username := ctx.Request.FormValue("userName")

	if username == "" {
		return
	}

	// 调用后台服务
	rsp, err := serviceClient.QueryUserByName(context.TODO(), &us.Request{
		UserName: username,
	})
	if err != nil {
		resp.Response(http.StatusInternalServerError, err_code.ERROR, err)
		return
	}

	// 生成token
	rsp2, err := authClient.MakeAccessToken(context.TODO(), &au.Request{
		UserId:   rsp.User.Id,
		UserName: rsp.User.Name,
	})

	if err != nil {
		log.Info("[Login] 创建token失败，err：%s", zap.Any("err", err))
		resp.Response(http.StatusInternalServerError, err_code.ERROR, err)
		return
	}

	log.Info("[Login] token %s", zap.Any("Token", rsp2.Token))
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
		log.Info("token获取失败")
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
