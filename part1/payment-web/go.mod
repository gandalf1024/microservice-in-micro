module microservice-in-micro/part1/payment-web

go 1.13

replace bac => ../bac

replace auth => ../auth

replace plugins => ../plugins

require (
	auth v0.0.0-00010101000000-000000000000
	bac v0.0.0-00010101000000-000000000000
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.17.1
	github.com/micro/go-plugins v1.5.1
	plugins v0.0.0-00010101000000-000000000000
)
