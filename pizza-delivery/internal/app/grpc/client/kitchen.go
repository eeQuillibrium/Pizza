package grpcclient

import (
	"context"

	grpc_orders "github.com/eeQuillibrium/protos/gen/go/orders"
	"google.golang.org/grpc"
)

type KitchenClient struct {
	grpcConn   *grpc.ClientConn
	grpcClient grpc_orders.OrderingClient
}

func NewKitchenClient(
	grpcConn *grpc.ClientConn,
	grpcClient grpc_orders.OrderingClient,
) *KitchenClient {
	return &KitchenClient{
		grpcConn:   grpcConn,
		grpcClient: grpcClient,
	}
}
func (c *KitchenClient) SendOrder(
	ctx context.Context,
	in *grpc_orders.SendOrderReq,
) (*grpc_orders.EmptyOrderResp, error) {
	return c.grpcClient.SendOrder(ctx, in)
}
