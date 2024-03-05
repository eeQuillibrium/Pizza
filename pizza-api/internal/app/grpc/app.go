package grpcapp

import (
	"context"
	"fmt"
	"log"

	grpcauth "github.com/eeQuillibrium/pizza-api/internal/app/grpc/auth"
	grpckitchen "github.com/eeQuillibrium/pizza-api/internal/app/grpc/kitchen"
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
	SendMessage(
		ctx context.Context,
		in *nikita_kitchen1.SendOrderReq,
	)
}

type GRPCApp struct {
	Auth    Auth
	Kitchen Kitchen
	//other grpc apps
}

func New(
	authport int,
	kitchenport int,
) *GRPCApp {
	log.Print("trying to set connection with authgrpc server...")

	authconn := setConn(authport)
	auth := grpcauth.New(authport, authconn)

	kitchenconn := setConn(kitchenport)
	kitchen := grpckitchen.New(kitchenport, kitchenconn)

	return &GRPCApp{
		auth,
		kitchen,
	}
}
func setConn(port int) *grpc.ClientConn {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%v", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect with auth service: %v", err)
	}
	return conn
}
func (a *GRPCApp) Stop() {

}
