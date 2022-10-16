package order

import (
	"context"
	"fmt"

	v1 "github.com/blck-snwmn/gomodmonorpc/modules/proto/gen/grpc/order/v1"
)

type (
	Input  struct{}
	Output struct{}
)

type Order struct{}

func (oh *Order) Do(ctx context.Context, input Input) (Output, error) {
	return Output{}, nil
}

type OrderService struct {
	v1.UnimplementedOrderServiceServer
}

func (os *OrderService) Order(context.Context, *v1.OrderRequest) (*v1.OrderResponse, error) {
	fmt.Println("call: Order")
	return &v1.OrderResponse{
		Result: "ok",
	}, nil
}
