package client

import (
	"context"

	grpc_orders "github.com/eeQuillibrium/protos/gen/go/orders"
	"google.golang.org/grpc"
)

type KitchenClient struct {
	gRPCClient grpc_orders.OrderingClient
	conn       *grpc.ClientConn
	port       int
}

func NewKitchen(
	port int,
	conn *grpc.ClientConn,
) *KitchenClient {
	gRPCClient := grpc_orders.NewOrderingClient(conn)
	return &KitchenClient{
		gRPCClient,
		conn,
		port,
	}
}

func (c *KitchenClient) SendOrder(
	ctx context.Context,
	in *grpc_orders.SendOrderReq,
) (*grpc_orders.EmptyOrderResp, error) {
	return c.gRPCClient.SendOrder(ctx, in)
}

func (c *KitchenClient) Stop() {
	c.conn.Close()
}
