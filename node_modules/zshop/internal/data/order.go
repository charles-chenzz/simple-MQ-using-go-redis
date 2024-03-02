package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "zshop/api/order/v1"
	"zshop/internal/biz"
	"zshop/internal/types"
)

type orderRepo struct {
	data *Data
	log  *log.Helper
}

func NewOrderRepo(data *Data, logger log.Logger) biz.OrderRepo {
	return &orderRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *orderRepo) CreateOrder(ctx context.Context, req *v1.CreateOrderRequest) (rsp *v1.CreateOrderReply, err error) {
	r.log.WithContext(ctx).Infof("creating order req:%v", req)
	db := r.data.db

	u := types.User{
		UserId:   req.GetUser().GetUserId(),
		UserName: req.GetUser().GetUserName(),
	}

	sql := "place your sql here"

	result, err := db.Exec(sql)
	if err != nil {
		log.Errorf("insert/update/delete error:%v", err)
		return nil, err
	}

	count, _ := result.RowsAffected()
	log.Debugf("result:%v", count)

	orderField := req.GetOrderInformation().GetOrderField()
	shipping := req.GetOrderInformation().GetShippingAddress()
	// todo graceful solution
	order := types.Order{
		OrderId:       orderField.GetOrderId(),
		TransactionId: orderField.GetTransactionId(),
		ProductId:     orderField.GetProductId(),
		ProductType:   orderField.GetProductType(),
		Quantity:      orderField.GetQuantity(),
		Size:          orderField.GetSize(),
		Color:         orderField.GetColor(),
		Status:        0,
		RetryTime:     0,
	}

	ship := types.Shipping{
		Email:             shipping.GetEmail(),
		Address:           shipping.GetAddress(),
		FirstName:         shipping.GetFirstName(),
		LastName:          shipping.GetLastName(),
		ApartmentSuiteEtc: shipping.GetApartmentSuiteEtc(),
		City:              shipping.GetCity(),
		State:             shipping.GetState(),
		ZipCode:           shipping.GetZipCode(),
		Phone:             shipping.GetPhone(),
	}

	err = db.Select(&u, "place your query text here")
	if err != nil {
		log.Errorf("error:%v", err)
		return nil, err
	}

	err = db.Select(&order, "place your query text here")
	if err != nil {
		log.Errorf("error:%v", err)
		return nil, err
	}

	err = db.Select(&ship, "place your query text here")
	if err != nil {
		log.Errorf("error:%v", err)
		return nil, err
	}

	//todo mq publish for other service
	return rsp, err
}
