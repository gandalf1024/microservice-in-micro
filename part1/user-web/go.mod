module microservice-in-micro/part1/user-web

go 1.13

require (
	auth v0.0.0-00010101000000-000000000000
	bac v0.0.0-00010101000000-000000000000
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/gin-gonic/gin v1.5.0
	github.com/golang/protobuf v1.3.2
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.17.1
	github.com/micro/go-plugins v1.5.1
	github.com/opentracing/opentracing-go v1.1.0
	go.uber.org/zap v1.12.0
	plugins v0.0.0-00010101000000-000000000000
	user_srv v0.0.0-00010101000000-000000000000
)

replace bac => ../bac

replace user_srv => ../user-srv

replace auth => ../auth

replace plugins => ../plugins
