package main

import (
	"checkout-service/internal/service" // 引入服务逻辑
	"checkout-service/proto"            // 引入生成的Proto文件
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// 连接到PaymentService
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to payment service: %v", err)
	}
	defer conn.Close()

	paymentClient := proto.NewPaymentServiceClient(conn)

	// 创建CheckoutService实例
	checkoutService := service.NewCheckoutService(paymentClient)

	// 启动gRPC服务器
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// 注册CheckoutService
	proto.RegisterCheckoutServiceServer(grpcServer, checkoutService)

	// 启动服务器
	fmt.Println("Checkout service is running on port 50052")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
