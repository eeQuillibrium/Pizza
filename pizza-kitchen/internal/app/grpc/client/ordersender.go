package grpcclient

import (
	"context"

	grpc_orders "github.com/eeQuillibrium/protos/gen/go/orders"
	"google.golang.org/grpc"
)

type OrderSenderClient struct {
	grpcClient grpc_orders.OrderingClient
}

func NewOS(grpcConn *grpc.ClientConn) *OrderSenderClient {
	grpcClient := grpc_orders.NewOrderingClient(grpcConn)
	return &OrderSenderClient{
		grpcClient: grpcClient,
	}
}

func (g *OrderSenderClient) SendOrder(
	ctx context.Context,
	in *grpc_orders.SendOrderReq,
) (*grpc_orders.EmptyOrderResp, error) {
	return g.grpcClient.SendOrder(ctx, in)
}
