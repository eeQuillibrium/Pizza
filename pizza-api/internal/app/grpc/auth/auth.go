package grpcauth

import (
	"context"
	"log"

	nikita_auth1 "github.com/eeQuillibrium/protos/proto/gen/go/auth"
	"google.golang.org/grpc"
)

const (
	emptyInt = 0
)

type GRPCAuth struct {
	gRPCClient *nikita_auth1.AuthClient
	conn       *grpc.ClientConn
	port       string
}

func New(
	port string,
	conn *grpc.ClientConn,
) *GRPCAuth {
	gRPCClient := nikita_auth1.NewAuthClient(conn)
	return &GRPCAuth{
		&gRPCClient,
		conn,
		port,
	}
}

func (g *GRPCAuth) Register(
	ctx context.Context,
	in *nikita_auth1.RegRequest,
) (int64, error) {

	r, err := (*g.gRPCClient).Register(ctx, in)
	if err != nil {
		log.Fatalf("client grpc in auth(Register): %v", err)
		return emptyInt, err
	}

	return r.GetUserId(), nil
}
func (g *GRPCAuth) Login(
	ctx context.Context,
	in *nikita_auth1.LoginRequest,
) (string, error) {
	return "", nil
}
func (g *GRPCAuth) IsAdmin(
	ctx context.Context,
	in *nikita_auth1.IsAdminRequest,
) (bool, error) {
	return true, nil
}

func (g *GRPCAuth) Stop() {
	g.conn.Close()
}
