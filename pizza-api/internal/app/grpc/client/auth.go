package client

import (
	"context"

	nikita_auth1 "github.com/eeQuillibrium/protos/gen/go/auth"
	"google.golang.org/grpc"
)

const (
	emptyInt = 0
)

type AuthClient struct {
	gRPCClient nikita_auth1.AuthClient
	conn       *grpc.ClientConn
	port       int
}

func NewAuth(
	port int,
	conn *grpc.ClientConn,
) *AuthClient {
	gRPCClient := nikita_auth1.NewAuthClient(conn)
	return &AuthClient{
		gRPCClient,
		conn,
		port,
	}
}

func (g *AuthClient) Register(
	ctx context.Context,
	in *nikita_auth1.RegRequest,
) (int64, error) {
	r, err := g.gRPCClient.Register(ctx, in)
	if err != nil {
		return emptyInt, err
	}
	return r.GetUserId(), nil
}
func (g *AuthClient) Login(
	ctx context.Context,
	in *nikita_auth1.LoginRequest,
) (string, error) {
	r, err := g.gRPCClient.Login(ctx, in)
	if err != nil {
		return "", err
	}

	return r.GetToken(), nil
}
func (g *AuthClient) IsAdmin(
	ctx context.Context,
	in *nikita_auth1.IsAdminRequest,
) (bool, error) {
	return true, nil
}

func (g *AuthClient) Stop() {
	g.conn.Close()
}