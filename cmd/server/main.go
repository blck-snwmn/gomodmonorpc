package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/blck-snwmn/gomodmonorpc/modules/invoker"
	"github.com/blck-snwmn/gomodmonorpc/modules/order"
	"github.com/blck-snwmn/gomodmonorpc/modules/payment"
	myec "github.com/blck-snwmn/gomodmonorpc/modules/proto/gen/grpc/myec/v1"
	orderv1 "github.com/blck-snwmn/gomodmonorpc/modules/proto/gen/grpc/order/v1"
	paymentv1 "github.com/blck-snwmn/gomodmonorpc/modules/proto/gen/grpc/payment/v1"
	"google.golang.org/grpc"
)

type invokerx struct {
	psv paymentv1.PaymentServiceServer
	osv orderv1.OrderServiceServer
	v   *invoker.Store
}

func (c *invokerx) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, options ...grpc.CallOption) error {
	fmt.Printf("[%s]call\n", method)
	defer fmt.Printf("[%s]end\n", method)

	f, err := c.v.Invokef(method)
	if err != nil {
		return err
	}
	defer fmt.Printf("Invoke: %+[1]v, %[1]T\n", reply)
	return f(ctx, method, args, reply, options...)
}

func (*invokerx) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	panic("unimplemented")
}

var _ grpc.ClientConnInterface = (*invokerx)(nil)

type server struct {
	myec.UnimplementedMyECServiceServer
	pc paymentv1.PaymentServiceClient
	oc orderv1.OrderServiceClient
}

func NewServer() *server {
	psv := &payment.PaymentService{}
	osv := &order.OrderService{}
	s := invoker.NewStore()

	s.RegisterByServiceName(orderv1.OrderService_ServiceDesc.ServiceName, osv.Invoke)
	s.RegisterByServiceName(paymentv1.PaymentService_ServiceDesc.ServiceName, psv.Invoke)

	iv := &invokerx{
		psv: psv,
		osv: osv,
		v:   s,
	}
	return &server{
		pc: paymentv1.NewPaymentServiceClient(iv),
		oc: orderv1.NewOrderServiceClient(iv),
	}
}

func (s *server) Order(ctx context.Context, req *myec.OrderRequest) (*myec.OrderResponse, error) {
	resp, err := s.oc.Order(ctx, &orderv1.OrderRequest{})
	if err != nil {
		return nil, err
	}
	fmt.Printf("Order: %+[1]v, %[1]T, Result=%q\n", resp, resp.Result)
	return &myec.OrderResponse{}, nil
}

func (s *server) Payment(ctx context.Context, req *myec.OrderRequest) (*myec.OrderResponse, error) {
	resp, err := s.pc.Pay(ctx, &paymentv1.PayRequest{})
	if err != nil {
		return nil, err
	}
	fmt.Printf("Payment: %+[1]v, %[1]T, Result=%q\n", resp, resp.Result)
	return &myec.OrderResponse{}, nil
}

func main() {
	listenOn := "127.0.0.1:8080"
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		fmt.Printf("failed to listen on %s: %+v", listenOn, err)
		return
	}

	sv := grpc.NewServer()
	myec.RegisterMyECServiceServer(sv, NewServer())
	log.Println("Listening on", listenOn)
	if err := sv.Serve(listener); err != nil {
		fmt.Printf("failed to serve gRPC server: %+v", err)
		return
	}
}
