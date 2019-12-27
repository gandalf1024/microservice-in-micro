package main

import (
	"bac"
	"bac/common"
	"bac/config"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/cli"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/config/source/grpc"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/opentracing/opentracing-go"
	"microservice-in-micro/part1/user-web/handler"
	"microservice-in-micro/part1/user-web/routers"
	"net"
	"net/http"
	"plugins/breaker"
	tracer "plugins/tracer/jaeger"
	"plugins/tracer/opentracing/std2micro"
	"time"
)

var (
	appName = "user_web"
	cfg     = &userCfg{}
)

type userCfg struct {
	common.AppCfg
}

func main() {
	// 初始化配置
	initCfg()

	// 使用consul注册
	micReg := etcdv3.NewRegistry(registryOptions)

	// 创建新服务
	service := web.NewService(
		// 后面两个web，第一个是指是web类型的服务，第二个是服务自身的名字
		web.Name(cfg.Name),
		web.Version(cfg.Version),
		web.RegisterTTL(time.Second*15),      //健康检查
		web.RegisterInterval(time.Second*10), //健康检查
		web.Registry(micReg),
		web.Address(cfg.Addr()),
	)

	t, io, err := tracer.NewTracer(cfg.Name, "192.168.59.137:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// 初始化服务
	if err := service.Init(
		web.Action(
			func(c *cli.Context) {
				// 初始化handler
				handler.Init()
			}),
	); err != nil {
		log.Fatal(err)
	}

	//std2micro.TracerWrapper(breaker.BreakerWrapper(routers.InitRouter())) 添加链路追踪
	service.Handle("/", std2micro.TracerWrapper(breaker.BreakerWrapper(routers.InitRouter())))

	//docker run --name hystrix-dashboard -d -p 8081:9002 mlabouardy/hystrix-dashboard:latest
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	//给 192.168.59.137 容器提供数据
	go func() {
		err := http.ListenAndServe(net.JoinHostPort("0.0.0.0", "9981"), hystrixStreamHandler)
		if err != nil {
			fmt.Println("1111111111--------------->>>")
		} else {
			fmt.Println("2222222222--------------->>>")
		}
	}()

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := &common.Etcd{}
	err := config.C().App("etcd", etcdCfg)
	if err != nil {
		panic(err)
	}
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.Host, etcdCfg.Port)}
}

func initCfg() {
	source := grpc.NewSource(
		grpc.WithAddress("127.0.0.1:9600"),
		grpc.WithPath("micro"),
	)

	bac.Init(config.WithSource(source))

	err := config.C().App(appName, cfg)
	if err != nil {
		panic(err)
	}

	log.Logf("[initCfg] 配置，cfg：%v", cfg)

	return
}
