package grpcapp

import (
	"fmt"
	"log"
	"net"

	"github.com/eeQuillibrium/pizza-kitchen/internal/app/grpc/server"
	"github.com/eeQuillibrium/pizza-kitchen/internal/service"
	"google.golang.org/grpc"
)

type GRPCApp struct {
	portApi     int
	grpcServAPI *grpc.Server
}

func New(
	portApi int,
	kAPIService *service.Service,
	//grpcPortDel int,
) *GRPCApp {

	grpcServAPI := grpc.NewServer()
	server.Register(grpcServAPI, kAPIService)

	return &GRPCApp{
		portApi:     portApi,
		grpcServAPI: grpcServAPI,
	}
}

func (g *GRPCApp) Run() {
	log.Print("try to run grpc kitchenapi serv...")

	lst, err := net.Listen("tcp", fmt.Sprintf(":%d", g.portApi))

	if err != nil {
		log.Fatalf("listen was dropped")
	}

	if err := g.grpcServAPI.Serve(lst); err != nil {
		log.Fatalf("serving was dropped")
	}
}

func (a *GRPCApp) Stop() {
	log.Printf("Stopping gRPC server %v...", a.portApi)
	a.grpcServAPI.GracefulStop()
}
