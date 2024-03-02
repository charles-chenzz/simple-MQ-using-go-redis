package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "zshop/api/order/v1"
	"zshop/internal/types"
)

type OrderRepo interface {
	CreateOrder(ctx context.Context, order *v1.CreateOrderRequest) (rsp *v1.CreateOrderReply, err error)
}

type OrderUserCase struct {
	repo OrderRepo
	log  *log.Helper
}

func NewOrderUserCase(repo OrderRepo, logger log.Logger) *OrderUserCase {
	return &OrderUserCase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *OrderUserCase) CreateOrder(ctx context.Context, req *v1.CreateOrderRequest) (rsp *v1.CreateOrderReply, err error) {
	uc.log.WithContext(ctx).Infof("biz creating order, parameter:%v", req)

	// todo generate from idmaker base on config center
	req.OrderInformation.OrderField.OrderId = 10001
	req.OrderInformation.OrderField.TransactionId = 10086

	// biz -> data
	rsp, err = uc.repo.CreateOrder(ctx, req)
	if err != nil {
		rsp.ErrMessage = "create order error"
		rsp.ErrCode = types.DataCreateError
		log.Errorf("biz to data error:%v err_code:%v err_msg:%v", req, rsp.ErrCode, rsp.ErrMessage)
		return rsp, err
	}

	return rsp, err
}
