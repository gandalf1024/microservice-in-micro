module microservice-in-micro/part1/plugins

go 1.13

replace basic => ../basic

replace plugins => ../plugins

require (
	basic v0.0.0-00010101000000-000000000000
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/go-redis/redis v6.15.6+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/google/uuid v1.1.1
	github.com/gorilla/sessions v1.2.0
	github.com/micro/go-micro v1.17.1
	github.com/opentracing/opentracing-go v1.1.0
	github.com/uber/jaeger-client-go v2.21.1+incompatible
	go.uber.org/zap v1.12.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	plugins v0.0.0-00010101000000-000000000000
)
