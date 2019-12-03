module microservice-in-micro/part1/plugins

go 1.13

replace bac => ../bac

require (
	bac v0.0.0-00010101000000-000000000000
	github.com/go-redis/redis v6.15.6+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/google/uuid v1.1.1
	github.com/gorilla/sessions v1.2.0
	github.com/micro/go-micro v1.17.1
	go.uber.org/zap v1.12.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)
