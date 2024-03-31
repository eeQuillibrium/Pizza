package grpcclient

import (
	"context"

	grpc_orders "github.com/eeQuillibrium/protos/gen/go/orders"
	"google.golang.org/grpc"
)

type APIGatewayClient struct {
	grpcClient grpc_orders.OrderingClient
	grpcConn   *grpc.ClientConn
}

func NewAPIClient(
	grpcClient grpc_orders.OrderingClient,
	grpcConn *grpc.ClientConn,
) *APIGatewayClient {
	return &APIGatewayClient{
		grpcClient: grpcClient,
	}
}

func (c *APIGatewayClient) SendOrder(
	ctx context.Context,
	in *grpc_orders.SendOrderReq,
) (*grpc_orders.EmptyOrderResp, error) {
	return c.grpcClient.SendOrder(ctx, in)
}


