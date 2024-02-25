package grpcapp

import (
	"context"
	"fmt"
	"log"

	grpcauth "github.com/eeQuillibrium/pizza-api/internal/app/grpc/auth"
	nikita_auth1 "github.com/eeQuillibrium/protos/proto/gen/go/auth"
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

type GRPCApp struct {
	Auth Auth 
	//other grpc apps
}

func New(port string) *GRPCApp {
	log.Print("trying to set connection with authgrpc server...")

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%v", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect with auth service: %v", err)
	}
	auth := grpcauth.New(port, conn)

	return &GRPCApp{
		auth,
	}
}

func (a *GRPCApp) Stop() {

}
