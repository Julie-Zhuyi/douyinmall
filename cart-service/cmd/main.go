package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Julie-Zhuyi/douyinmall/cart-service/internal/service"
	pb "github.com/Julie-Zhuyi/douyinmall/cart-service/proto"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterCartServiceServer(server, &service.CartServiceImpl{})

	fmt.Println("Cart Service is running on port 50051...")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
