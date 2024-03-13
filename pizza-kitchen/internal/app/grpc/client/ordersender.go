package grpcclient

import (
	"context"
	"log"

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
	log.Printf("trying to proceed by ordersender")

	r, err := g.grpcClientAPI.SendOrder(ctx, in)

	if err != nil {
		log.Fatalf("sendOrder error: %v", err)
	}

	return r, err
}
