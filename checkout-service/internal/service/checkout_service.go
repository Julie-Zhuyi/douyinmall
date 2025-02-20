package service

import (
	"checkout-service/proto" // 引入生成的 checkout.proto 和 payment.proto 的Go代码
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CheckoutService struct {
	paymentClient proto.PaymentServiceClient // 引入 PaymentService 客户端
}

func NewCheckoutService(paymentClient proto.PaymentServiceClient) *CheckoutService {
	return &CheckoutService{
		paymentClient: paymentClient,
	}
}

func (s *CheckoutService) Checkout(ctx context.Context, req *proto.CheckoutReq) (*proto.CheckoutResp, error) {
	// 模拟订单ID生成
	orderID := fmt.Sprintf("ORD%v", time.Now().Unix())

	// 模拟支付请求
	chargeReq := &proto.ChargeReq{
		Amount:     100.0, // 假设金额为100
		OrderId:    orderID,
		UserId:     req.UserId,
		CreditCard: req.CreditCard,
	}

	// 调用 PaymentService 的 Charge 接口
	chargeResp, err := s.paymentClient.Charge(ctx, chargeReq)
	if err != nil {
		// 错误处理
		if status.Code(err) == codes.InvalidArgument {
			return nil, fmt.Errorf("invalid payment details: %v", err)
		}
		return nil, fmt.Errorf("failed to process payment: %v", err)
	}

	// 返回结算成功的响应
	return &proto.CheckoutResp{
		OrderId:       orderID,
		TransactionId: chargeResp.TransactionId,
	}, nil
}
