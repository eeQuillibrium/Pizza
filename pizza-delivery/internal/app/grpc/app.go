package grpcapp

import (
	"context"
	"fmt"
	"net"

	grpcclient "github.com/eeQuillibrium/pizza-delivery/internal/app/grpc/client"
	grpcserver "github.com/eeQuillibrium/pizza-delivery/internal/app/grpc/server"
	"github.com/eeQuillibrium/pizza-delivery/internal/logger"
	"github.com/eeQuillibrium/pizza-delivery/internal/service"
	grpc_orders "github.com/eeQuillibrium/protos/gen/go/orders"
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
	log           *logger.Logger
	GatewayClient OrderSender
	KitchenClient OrderSender
	OrderServer   *grpc.Server
}

func New(
	log *logger.Logger,
	orderProvider service.OrderProvider,
	gatewayPort int,
	kitchenPort int,
) *GRPCApp {
	gatewayConn := setConn(log, gatewayPort)
	gatewayClient := grpcclient.NewAPIClient(grpc_orders.NewOrderingClient(gatewayConn), gatewayConn)

	kitchenConn := setConn(log, kitchenPort)
	kitchenClient := grpcclient.NewKitchenClient(kitchenConn, grpc_orders.NewOrderingClient(kitchenConn))

	orderServer := grpc.NewServer()
	grpcserver.Register(orderServer, orderProvider)

	return &GRPCApp{
		log:           log,
		GatewayClient: gatewayClient,
		KitchenClient: kitchenClient,
		OrderServer:   orderServer,
	}
}

func setConn(
	log *logger.Logger,
	port int,
) *grpc.ClientConn {
	conn, err := grpc.Dial(fmt.Sprintf(":%d", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.SugaredLogger.Infof("conn failed on port %d, err: %w", port, err)
	}
	return conn
}

func (a *GRPCApp) Run(gatewayServerPort int) {

	lst, err := net.Listen("tcp", fmt.Sprintf(":%d", gatewayServerPort))
	if err != nil {
		a.log.SugaredLogger.Infof("listen problem : %w", err)
	}

	if err := a.OrderServer.Serve(lst); err != nil {
		a.log.SugaredLogger.Infof("grpc serving problem : %w", err)
	}
}

func (a *GRPCApp) Stop() {
	a.log.SugaredLogger.Info("stopping grpc server")
	a.OrderServer.GracefulStop()
}
