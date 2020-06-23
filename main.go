package main

import (
	proto "GoMicroTest/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro"
)

type GreeterServiceHandler struct{}

func (g *GreeterServiceHandler) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = " 我是 server，卑微的 client " + req.Name + "，有什麼事嗎!!!!"
	return nil
}

func main() {
	// 创建新的服务
	service := micro.NewService(
		micro.Name("Greeter"),
	)

	// 初始化，会解析命令行参数
	service.Init()

	// 注册处理器，调用 Greeter 服务接口处理请求
	proto.RegisterGreeterHandler(service.Server(), new(GreeterServiceHandler))

	// 启动服务
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
