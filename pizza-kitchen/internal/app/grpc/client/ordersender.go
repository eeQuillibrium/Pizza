package grpcclient

import (
	"context"

	grpc_orders "github.com/eeQuillibrium/protos/gen/go/orders"
	"google.golang.org/grpc"
)

type OrderSenderClient struct {
	grpcClientAPI grpc_orders.OrderingClient
	grpcConnAPI   *grpc.ClientConn
}

func NewOS(grpcConnAPI *grpc.ClientConn) *OrderSenderClient {
	grpcClientAPI := grpc_orders.NewOrderingClient(grpcConnAPI)
	return &OrderSenderClient{
		grpcClientAPI: grpcClientAPI,
		grpcConnAPI:   grpcConnAPI,
	}
}

func (g *OrderSenderClient) SendOrder(
	ctx context.Context,
	in *grpc_orders.SendOrderReq,
) (*grpc_orders.EmptyOrderResp, error) {
	return g.grpcClientAPI.SendOrder(ctx, in)
}
