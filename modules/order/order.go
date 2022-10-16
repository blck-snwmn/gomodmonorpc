package order

import (
	"context"
	"fmt"

	"github.com/blck-snwmn/gomodmonorpc/modules/invoker"
	v1 "github.com/blck-snwmn/gomodmonorpc/modules/proto/gen/grpc/order/v1"
	"google.golang.org/grpc"
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

func (os *OrderService) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, _ ...grpc.CallOption) error {
	switch req := args.(type) {
	case *v1.OrderRequest:
		resp, err := os.Order(ctx, req)
		if err != nil {
			return err
		}
		invoker.Cp(reply.(*v1.OrderResponse), resp)
		return nil
	default:
		panic("invalid args")
	}
}

func (os *OrderService) Order(context.Context, *v1.OrderRequest) (*v1.OrderResponse, error) {
	fmt.Println("call: Order")
	return &v1.OrderResponse{
		Result: "ok",
	}, nil
}
