package service

import (
	"context"
	"fmt"

	pb "github.com/Julie-Zhuyi/douyinmall/cart-service/proto"
)

// CartServiceImpl 购物车服务实现
type CartServiceImpl struct {
	pb.UnimplementedCartServiceServer
}

// AddItem 添加商品到购物车
func (s *CartServiceImpl) AddItem(ctx context.Context, req *pb.AddItemReq) (*pb.AddItemResp, error) {
	fmt.Printf("User %d added product %d (qty: %d)\n", req.UserId, req.Item.ProductId, req.Item.Quantity)
	return &pb.AddItemResp{}, nil
}

// GetCart 获取购物车
func (s *CartServiceImpl) GetCart(ctx context.Context, req *pb.GetCartReq) (*pb.GetCartResp, error) {
	return &pb.GetCartResp{
		UserId: req.UserId,
		Items: []*pb.CartItem{
			{ProductId: 1, Quantity: 2},
			{ProductId: 2, Quantity: 1},
		},
	}, nil
}

// EmptyCart 清空购物车
func (s *CartServiceImpl) EmptyCart(ctx context.Context, req *pb.EmptyCartReq) (*pb.EmptyCartResp, error) {
	fmt.Printf("User %d's cart emptied\n", req.UserId)
	return &pb.EmptyCartResp{}, nil
}
