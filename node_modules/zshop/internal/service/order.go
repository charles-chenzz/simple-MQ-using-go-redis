package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"zshop/internal/biz"
	"zshop/internal/types"

	pb "zshop/api/order/v1"
)

type OrderService struct {
	pb.UnimplementedOrderServer
	uc *biz.OrderUserCase
}

func NewOrderService(uc *biz.OrderUserCase) *OrderService {
	return &OrderService{uc: uc}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (rsp *pb.CreateOrderReply, err error) {
	log.Debugf("req:%v", req)

	valid := types.ValidateOrderField(req)
	if !valid {
		rsp.ErrCode = types.ValidatingError
		rsp.ErrMessage = "error"
		log.Errorf("order parameter unfulfilled, please try again later req:%v err_code:%v err_msg:%v", req, rsp.ErrCode, rsp.ErrMessage)
		return rsp, nil
	}

	rsp, err = s.uc.CreateOrder(ctx, req)
	if err != nil {
		log.Errorf("create order error, rsp:%v", rsp)
		return rsp, err
	}

	return rsp, nil
}
func (s *OrderService) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderReply, error) {
	return &pb.GetOrderReply{}, nil
}
