package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Julie-Zhuyi/douyinmall/order-service/internal/service"
	"github.com/Julie-Zhuyi/douyinmall/order-service/proto"

	"google.golang.org/grpc"
)

func main() {
	// 监听 50052 端口
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 创建gRPC服务器
	server := grpc.NewServer()

	// 注册OrderService服务
	proto.RegisterOrderServiceServer(server, &service.OrderService{})

	// 启动服务器
	fmt.Println("Order Service is running on port 50052")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
