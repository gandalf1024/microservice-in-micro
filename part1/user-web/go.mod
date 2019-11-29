module microservice-in-micro/part1/user-web

go 1.13

require (
	auth v0.0.0-00010101000000-000000000000
	bac v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.5.0
	github.com/golang/protobuf v1.3.2
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.17.1
	github.com/micro/go-plugins v1.5.1
	user_srv v0.0.0-00010101000000-000000000000
)

replace bac => ../bac

replace user_srv => ../user-srv

replace auth => ../auth
