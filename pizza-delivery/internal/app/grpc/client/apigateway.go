package grpcclient

import (
	"context"

	grpc_orders "github.com/eeQuillibrium/protos/gen/go/orders"
)

type APIGatewayClient struct {
	Client grpc_orders.OrderingClient
}

func NewAPIClient(
	Client grpc_orders.OrderingClient,
) *APIGatewayClient {
	return &APIGatewayClient{
		Client: Client,
	}
}

func (c *APIGatewayClient) SendOrder(
	ctx context.Context,
	in *grpc_orders.SendOrderReq,
) (*grpc_orders.EmptyOrderResp, error) {
	return c.Client.SendOrder(ctx, in)
}
