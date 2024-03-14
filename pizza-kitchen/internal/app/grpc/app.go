package grpcapp

import (
	"context"
	"fmt"
	"net"

	grpc_orders "github.com/eeQuillibrium/protos/gen/go/orders"

	grpcclient "github.com/eeQuillibrium/pizza-kitchen/internal/app/grpc/client"
	"github.com/eeQuillibrium/pizza-kitchen/internal/app/grpc/server"
	"github.com/eeQuillibrium/pizza-kitchen/internal/service"
	"github.com/eeQuillibrium/pizza-kitchen/internal/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OrderSender interface {
	SendOrder(
		ctx context.Context,
		in *grpc_orders.SendOrderReq,
	) (*grpc_orders.EmptyOrderResp, error)
}

type GRPCApp struct {
	log *logger.Logger
	portServAPI int
	grpcServAPI *grpc.Server
	APIOrderSender OrderSender
}

func New(
	log *logger.Logger,
	portClientAPI int,
	portServAPI int,
	kAPIService *service.Service,
	//grpcPortDel int,
) *GRPCApp {
	orderConn := setConn(log, portClientAPI)
	apiClient := grpcclient.NewOS(orderConn)

	grpcServAPI := grpc.NewServer()
	server.Register(grpcServAPI, kAPIService)

	return &GRPCApp{
		log: log,
		portServAPI: portServAPI,
		grpcServAPI: grpcServAPI,
		APIOrderSender: apiClient,
	}
}

func (g *GRPCApp) Run() {

	lst, err := net.Listen("tcp", fmt.Sprintf(":%d", g.portServAPI))

	if err != nil {
		g.log.SugaredLogger.Info("listen was dropped")
	}

	if err := g.grpcServAPI.Serve(lst); err != nil {
		g.log.SugaredLogger.Info("serving was dropped")
	}

}

func (g *GRPCApp) Stop() {
	g.log.SugaredLogger.Infof("stopping gRPC server %v", g.portServAPI)
	g.grpcServAPI.GracefulStop()
}

func setConn(log *logger.Logger, port int) *grpc.ClientConn {
	log.SugaredLogger.Infof("try to set connection on port %d", port)

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.SugaredLogger.Infof("failed to connect with auth service: %w", err)
	}

	return conn
}
