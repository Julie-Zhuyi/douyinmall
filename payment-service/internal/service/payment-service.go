package service

import (
	"context"
	"fmt"
	"time"

	"github.com/Julie-Zhuyi/douyinmall/payment-service/db"
	"github.com/Julie-Zhuyi/douyinmall/payment-service/proto"
)

type PaymentService struct {
	proto.UnimplementedPaymentServiceServer
}

// Charge 处理支付请求
func (s *PaymentService) Charge(ctx context.Context, req *proto.ChargeReq) (*proto.ChargeResp, error) {
	// 模拟支付处理
	transactionID := fmt.Sprintf("TXN%v", time.Now().Unix())

	// 插入支付记录到数据库
	transaction := db.Transaction{
		OrderID:       req.OrderId,
		TransactionID: transactionID,
		Amount:        float64(req.Amount),
		UserID:        uint(req.UserId),
		CreditCardInfo: fmt.Sprintf("%s %d/%d",
			req.CreditCard.CreditCardNumber, req.CreditCard.CreditCardExpirationMonth, req.CreditCard.CreditCardExpirationYear),
	}

	if err := db.DB.Create(&transaction).Error; err != nil {
		return nil, fmt.Errorf("failed to record transaction: %v", err)
	}

	// 返回支付成功的 transaction ID
	return &proto.ChargeResp{
		TransactionId: transactionID,
	}, nil
}
