package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/blck-snwmn/gomodmonorpc/modules/order"
	"github.com/blck-snwmn/gomodmonorpc/modules/payment"
	"github.com/blck-snwmn/gomodmonorpc/modules/proto/gen/buf/proto/myec/v1"
	"google.golang.org/grpc"
)

type server struct {
	myec.UnimplementedMyECServiceServer
	o *order.Order
	p *payment.Payment
}

func (s *server) Order(context.Context, *myec.OrderRequest) (*myec.OrderResponse, error) {
	_, err := s.o.Do(context.Background(), order.Input{})
	if err != nil {
		return nil, err
	}
	return &myec.OrderResponse{}, nil
}

func (s *server) Payment(context.Context, *myec.OrderRequest) (*myec.OrderResponse, error) {
	_, err := s.p.Do(context.Background(), payment.Input{})
	if err != nil {
		return nil, err
	}
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
	myec.RegisterMyECServiceServer(sv, &server{})
	log.Println("Listening on", listenOn)
	if err := sv.Serve(listener); err != nil {
		fmt.Printf("failed to serve gRPC server: %+v", err)
		return
	}
}
