package res

import (
	"github.com/gin-gonic/gin"
	"microservice-in-micro/part1/user-web/gin/err_code"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  err_code.GetMsg(errCode),
		Data: data,
	})
	return
}
