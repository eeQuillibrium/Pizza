package grpcapp

import (
	"context"
	"fmt"
	"log"
	"net"

	grpcauth "github.com/eeQuillibrium/pizza-api/internal/app/grpc/auth"
	grpckitchen "github.com/eeQuillibrium/pizza-api/internal/app/grpc/kitchen"
	grpcserver "github.com/eeQuillibrium/pizza-api/internal/app/grpc/server"
	"github.com/eeQuillibrium/pizza-api/internal/service"
	nikita_auth1 "github.com/eeQuillibrium/protos/gen/go/auth"
	nikita_kitchen1 "github.com/eeQuillibrium/protos/gen/go/kitchen"
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
}

type Kitchen interface {
	SendOrder(
		ctx context.Context,
		in *nikita_kitchen1.SendOrderReq,
	) (*nikita_kitchen1.EmptyOrderResp, error)
}

type GRPCApp struct {
	Auth        Auth
	Kitchen     Kitchen
	OrderServer *grpc.Server
	//other grpc
}

func New(
	authport int,
	kitchenport int,
	kService service.OrderProvider,
) *GRPCApp {
	log.Print("trying to set connection with authgrpc server...")

	authconn := setConn(authport)
	auth := grpcauth.New(authport, authconn)
	log.Print("authgrpc connect successful!")

	log.Print("trying to set connection with kitchen server...")
	kitchenconn := setConn(kitchenport)
	kitchen := grpckitchen.New(kitchenport, kitchenconn)
	log.Print("kitchen connect successful!")

	serv := grpc.NewServer()
	grpcserver.Register(serv, kService)

	return &GRPCApp{
		Auth:        auth,
		Kitchen:     kitchen,
		OrderServer: serv,
	}
}

func setConn(port int) *grpc.ClientConn {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%v", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect with auth service: %v", err)
	}
	return conn
}

func (a *GRPCApp) Run(
	orderPort int,
) {
	log.Printf("try to run grpc kitchenapi serv on %s", fmt.Sprintf(":%d", orderPort))

	lst, err := net.Listen("tcp", fmt.Sprintf(":%d", orderPort))

	if err != nil {
		log.Fatalf("listen was dropped")
	}

	if err := a.OrderServer.Serve(lst); err != nil {
		log.Fatalf("serving was dropped")
	}
}

func (a *GRPCApp) Stop() {
	a.OrderServer.GracefulStop()
}
