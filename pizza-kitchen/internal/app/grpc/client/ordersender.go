package grpcclient

import (
	"context"
	"log"

	nikita_kitchen1 "github.com/eeQuillibrium/protos/gen/go/kitchen"
	"google.golang.org/grpc"
)

type OrderSenderClient struct {
	grpcClientAPI nikita_kitchen1.KitchenClient
	grpcConnAPI   *grpc.ClientConn
}

func NewOS(grpcConnAPI *grpc.ClientConn) *OrderSenderClient {
	grpcClientAPI := nikita_kitchen1.NewKitchenClient(grpcConnAPI)
	return &OrderSenderClient{
		grpcClientAPI: grpcClientAPI,
		grpcConnAPI:   grpcConnAPI,
	}
}

func (g *OrderSenderClient) SendOrder(
	ctx context.Context,
	in *nikita_kitchen1.SendOrderReq,
) (*nikita_kitchen1.EmptyOrderResp, error) {
	log.Printf("trying to proceed by ordersender")

	r, err := g.grpcClientAPI.SendOrder(ctx, in)

	if err != nil {
		log.Fatalf("sendOrder error: %v", err)
	}

	return r, err
}
