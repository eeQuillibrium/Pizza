package client

import (
	"context"

	grpc_orders "github.com/eeQuillibrium/protos/gen/go/orders"
	"google.golang.org/grpc"
)

type DeliveryClient struct {
	gRPCClient grpc_orders.OrderingClient
	conn       *grpc.ClientConn
	port       int
}

func NewDelivery(
	port int,
	conn *grpc.ClientConn,
) *DeliveryClient {
	gRPCClient := grpc_orders.NewOrderingClient(conn)
	return &DeliveryClient{
		gRPCClient,
		conn,
		port,
	}
}
func (c *DeliveryClient) SendOrder(
	ctx context.Context,
	in *grpc_orders.SendOrderReq,
) (*grpc_orders.EmptyOrderResp, error) {
	return c.gRPCClient.SendOrder(ctx, in)
}

func (c *DeliveryClient) Stop() {
	c.conn.Close()
}
