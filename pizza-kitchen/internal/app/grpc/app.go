package grpcapp

import (
	"context"
	"fmt"
	"log"
	"net"

	nikita_kitchen1 "github.com/eeQuillibrium/protos/gen/go/kitchen"

	"github.com/eeQuillibrium/pizza-kitchen/internal/app/grpc/ordersender"
	"github.com/eeQuillibrium/pizza-kitchen/internal/app/grpc/server"
	"github.com/eeQuillibrium/pizza-kitchen/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OrderSender interface {
	SendOrder(
		ctx context.Context,
		in *nikita_kitchen1.SendOrderReq,
	) (*nikita_kitchen1.EmptyOrderResp, error)
}

type GRPCApp struct {
	portServAPI int
	grpcServAPI *grpc.Server
	OrderSender
}

func New(
	portClientAPI int,
	portServAPI int,
	kAPIService *service.Service,
	//grpcPortDel int,
) *GRPCApp {
	orderConn := setConn(portClientAPI)
	orderSender := ordersender.New(orderConn)

	grpcServAPI := grpc.NewServer()
	server.Register(grpcServAPI, kAPIService)

	return &GRPCApp{
		portServAPI: portServAPI,
		grpcServAPI: grpcServAPI,
		OrderSender: orderSender,
	}
}

func (g *GRPCApp) Run() {

	log.Print("try to run grpc kitchenapi serv...")

	lst, err := net.Listen("tcp", fmt.Sprintf(":%d", g.portServAPI))

	if err != nil {
		log.Fatalf("listen was dropped")
	}

	if err := g.grpcServAPI.Serve(lst); err != nil {
		log.Fatalf("serving was dropped")
	}

}

func (a *GRPCApp) Stop() {
	log.Printf("Stopping gRPC server %v...", a.portServAPI)
	a.grpcServAPI.GracefulStop()
}

func setConn(port int) *grpc.ClientConn {
	log.Printf("try to set connection on port %d", port)
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%v", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect with auth service: %v", err)
	}
	return conn
}
