package main

import (
	"bac"
	"bac/config"
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/registry/etcdv3"
	"microservice-in-micro/part1/user-srv/handler"
	"microservice-in-micro/part1/user-srv/model"
	user "microservice-in-micro/part1/user-srv/proto/user"
	"time"
)

func main() {
	// 初始化配置、数据库等信息
	bac.Init()

	reg := etcdv3.NewRegistry(registryOptions)

	// 初始化服务
	service := micro.NewService(
		micro.Name("mu.micro.book.srv.user"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Registry(reg),
	)

	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) {
			// 初始化模型层
			model.Init()
			// 初始化handler
			handler.Init()
		}),
	)

	_ = user.RegisterUserHandler(service.Server(), new(handler.Service))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	consulCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", consulCfg.GetHost(), consulCfg.GetPort())}
}
