package types

import (
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	pb "zshop/api/order/v1"
)

func ValidateOrderField(req *pb.CreateOrderRequest) (result bool) {
	//bv := wrapperspb.BoolValue{}
	//fd := (&bv).ProtoReflect().Descriptor().Fields().ByName("order_information")
	//result := bv.ProtoReflect().Has(fd)

	orderField := req.GetOrderInformation().GetOrderField()
	shipping := req.GetOrderInformation().GetShippingAddress()

	result = true

	if orderField.GetProductId() < 0 || orderField.GetProductType() < 0 || orderField.GetQuantity() < 0 || len(orderField.GetSize()) < 0 || len(orderField.GetColor()) < 0 {
		result = false
	}

	if len(shipping.GetEmail()) < 0 || len(shipping.GetAddress()) < 0 || len(shipping.GetFirstName()) < 0 || len(shipping.GetLastName()) < 0 ||
		len(shipping.GetApartmentSuiteEtc()) < 0 || len(shipping.GetCity()) < 0 || shipping.State < 0 || shipping.GetZipCode() < 0 || shipping.GetPhone() < 0 {
		result = false
	}

	if req.GetUser().GetIsMember() && (req.GetUser().GetUserId() < 0 || len(req.GetUser().GetUserName()) < 0) {
		result = false
	}

	return result
}
