module microservice-in-micro/part1/orders-srv

go 1.13

replace bac => ../bac

replace inventory_srv => ../inventory-srv

replace payment_srv => ../payment-srv

require (
	bac v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/protobuf v1.3.2
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.17.1
	github.com/micro/go-plugins v1.5.1
	github.com/micro/protoc-gen-micro v1.0.0
	inventory_srv v0.0.0-00010101000000-000000000000
	payment_srv v0.0.0-00010101000000-000000000000
)
