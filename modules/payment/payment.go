package payment

import (
	"context"
	"fmt"

	v1 "github.com/blck-snwmn/gomodmonorpc/modules/proto/gen/grpc/payment/v1"
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

func (ps *PaymentService) Pay(context.Context, *v1.PayRequest) (*v1.PayResponse, error) {
	fmt.Println("call: Pay")
	return &v1.PayResponse{
		Result: "ok",
	}, nil
}
