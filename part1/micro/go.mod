module microservice-in-micro/part1/micro

go 1.13

replace plugins => ../plugins

replace basic => ../basic

require (
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/micro/go-plugins v1.5.1
	github.com/micro/micro v1.18.0
	github.com/opentracing/opentracing-go v1.1.0
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	plugins v0.0.0-00010101000000-000000000000
)
