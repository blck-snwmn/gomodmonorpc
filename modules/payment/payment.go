package payment

import (
	"context"
	"fmt"

	"github.com/blck-snwmn/gomodmonorpc/modules/invoker"
	v1 "github.com/blck-snwmn/gomodmonorpc/modules/proto/gen/grpc/payment/v1"
	"google.golang.org/grpc"
)

type (
	Input  struct{}
	Output struct{}
)

type Payment struct{}

func (oh *Payment) Do(ctx context.Context, input Input) (Output, error) {
	return Output{}, nil
}

type PaymentService struct {
	v1.UnimplementedPaymentServiceServer
}

func (ps *PaymentService) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, _ ...grpc.CallOption) error {
	switch req := args.(type) {
	case *v1.PayRequest:
		resp, err := ps.Pay(ctx, req)
		if err != nil {
			return err
		}
		invoker.Cp(reply.(*v1.PayResponse), resp)
		return nil
	default:
		panic("invalid args")
	}
}

func (ps *PaymentService) Pay(context.Context, *v1.PayRequest) (*v1.PayResponse, error) {
	fmt.Println("call: Pay")
	return &v1.PayResponse{
		Result: "ok",
	}, nil
}
