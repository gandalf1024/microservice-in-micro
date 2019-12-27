module microservice-in-micro/part1/orders-web

go 1.13

replace bac => ../bac

replace payment_srv => ../payment-srv

replace inventory_srv => ../inventory-srv

replace orders_srv => ../orders-srv

replace auth => ../auth

replace plugins => ../plugins

require (
	auth v0.0.0-00010101000000-000000000000
	bac v0.0.0-00010101000000-000000000000
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.17.1
	github.com/micro/go-plugins v1.5.1
	github.com/opentracing/opentracing-go v1.1.0
	inventory_srv v0.0.0-00010101000000-000000000000
	orders_srv v0.0.0-00010101000000-000000000000
	plugins v0.0.0-00010101000000-000000000000
)
