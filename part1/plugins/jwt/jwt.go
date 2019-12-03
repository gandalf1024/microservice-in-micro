package jwt

import "bac"

// Jwt 配置 接口
type Jwt struct {
	SecretKey string `json:"secretKey"`
}

// init 初始化Redis
func init() {
	bac.Register(initJwt)
}

func initJwt() {

}
