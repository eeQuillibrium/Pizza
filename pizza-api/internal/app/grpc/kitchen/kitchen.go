package grpckitchen

import (
	"context"

	nikita_kitchen1 "github.com/eeQuillibrium/protos/gen/go/kitchen"
	"google.golang.org/grpc"
)

type Kitchen struct {
	gRPCClient nikita_kitchen1.KitchenClient
	conn       *grpc.ClientConn
	port       int
}

func New(
	port int,
	conn *grpc.ClientConn,
) *Kitchen {
	gRPCClient := nikita_kitchen1.NewKitchenClient(conn)
	return &Kitchen{
		gRPCClient,
		conn,
		port,
	}
}

func (k *Kitchen) SendMessage(
	ctx context.Context,
	in *nikita_kitchen1.SendOrderReq,
)
