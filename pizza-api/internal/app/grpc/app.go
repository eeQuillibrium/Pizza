package grpcapp

import (
	"context"
	"fmt"
	"net"

	"github.com/eeQuillibrium/pizza-api/internal/app/grpc/client"
	grpcserver "github.com/eeQuillibrium/pizza-api/internal/app/grpc/server"
	"github.com/eeQuillibrium/pizza-api/internal/logger"
	"github.com/eeQuillibrium/pizza-api/internal/service"
	nikita_auth1 "github.com/eeQuillibrium/protos/gen/go/auth"
	grpc_orders "github.com/eeQuillibrium/protos/gen/go/orders"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// requesters; see (grpc/auth)
type Auth interface {
	Register(
		ctx context.Context,
		in *nikita_auth1.RegRequest,
	) (int64, error)
	Login(
		ctx context.Context,
		in *nikita_auth1.LoginRequest,
	) (string, error)
	IsAdmin(
		ctx context.Context,
		in *nikita_auth1.IsAdminRequest,
	) (bool, error)
	UserIdentify(
		ctx context.Context,
		in *nikita_auth1.IdentifyRequest,
	) (int, error)
}

type OrderSender interface {
	SendOrder(
		ctx context.Context,
		in *grpc_orders.SendOrderReq,
	) (*grpc_orders.EmptyOrderResp, error)
}

type GRPCApp struct {
	log         *logger.Logger
	serverPort  int
	OrderServer *grpc.Server
	Auth        Auth
	KitchenOS   OrderSender
	DeliveryOS  OrderSender
	//other grpc
}

func New(
	log *logger.Logger,
	serverPort int,
	authport int,
	kitchenport int,
	deliveryport int,
	kService service.OrderProvider,
) *GRPCApp {
	log.SugaredLogger.Info("trying to set connection with authgrpc server...")

	authconn := setConn(log, authport)
	auth := client.NewAuth(authport, authconn)
	log.SugaredLogger.Info("authgrpc connect successful!")

	log.SugaredLogger.Info("trying to set connection with kitchen server...")
	kitchenconn := setConn(log, kitchenport)
	kitchenOS := client.NewKitchen(kitchenport, kitchenconn)
	log.SugaredLogger.Info("kitchen connect successful!")

	log.SugaredLogger.Info("trying to set connection with delivery server...")
	deliveryConn := setConn(log, deliveryport)
	deliveryOS := client.NewDelivery(deliveryport, deliveryConn)
	log.SugaredLogger.Info("delivery connect successful!")

	serv := grpc.NewServer()
	grpcserver.Register(serv, kService)

	return &GRPCApp{
		log:         log,
		serverPort:  serverPort,
		OrderServer: serv,
		Auth:        auth,
		KitchenOS:   kitchenOS,
		DeliveryOS:  deliveryOS,
	}
}

func setConn(
	log *logger.Logger,
	port int,
) *grpc.ClientConn {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%v", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.SugaredLogger.Fatalf("failed to connect with auth service: %w", err)
	}
	return conn
}

func (a *GRPCApp) Run() {
	a.log.SugaredLogger.Infof("run grpc server on %d", a.serverPort)

	lst, err := net.Listen("tcp", fmt.Sprintf(":%d", a.serverPort))

	if err != nil {
		a.log.SugaredLogger.Fatal("listen was dropped")
	}

	if err := a.OrderServer.Serve(lst); err != nil {
		a.log.SugaredLogger.Fatal("serving was dropped")
	}

}

func (a *GRPCApp) Stop() {
	a.log.SugaredLogger.Infof("stopping grpc server")
	a.OrderServer.GracefulStop()
}
