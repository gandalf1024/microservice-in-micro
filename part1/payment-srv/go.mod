module microservice-in-micro/part1/payment-srv

go 1.13

replace basic => ../basic

replace plugins => ../plugins

require (
	basic v0.0.0-00010101000000-000000000000
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/protobuf v1.3.2
	github.com/google/uuid v1.1.1
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.17.1
	github.com/micro/go-plugins v1.5.1
	github.com/opentracing/opentracing-go v1.1.0
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	plugins v0.0.0-00010101000000-000000000000
)
