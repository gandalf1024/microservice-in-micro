module microservice-in-micro/part1/inventory-srv

go 1.13

require (
	bac v0.0.0-00010101000000-000000000000
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/protobuf v1.3.2
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.17.1
	github.com/micro/go-plugins v1.5.1
	github.com/opentracing/opentracing-go v1.1.0
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	plugins v0.0.0-00010101000000-000000000000
)

replace bac => ../bac

replace user_srv => ../user-srv

replace auth => ../auth

replace plugins => ../plugins
