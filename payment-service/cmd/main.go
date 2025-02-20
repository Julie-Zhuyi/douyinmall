package main

import (
	"fmt"
	"net"

	"github.com/Julie-Zhuyi/douyinmall/payment-service/db"
	"github.com/Julie-Zhuyi/douyinmall/payment-service/internal/service"
	"github.com/Julie-Zhuyi/douyinmall/payment-service/proto"
	"google.golang.org/grpc"
)

func main() {
	// 初始化数据库
	db.InitializeDB()

	// 启动 gRPC 服务器
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}

	grpcServer := grpc.NewServer()
	proto.RegisterPaymentServiceServer(grpcServer, &service.PaymentService{})
	fmt.Println("Payment service started at :50053")
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v\n", err)
	}
}
