package server

import (
	"context"
	"log"

	grpc_orders "github.com/eeQuillibrium/protos/gen/go/orders"
)

// orderprovider server handler implementation
func (s *serverAPI) SendOrder(
	ctx context.Context,
	in *grpc_orders.SendOrderReq,
) (*grpc_orders.EmptyOrderResp, error) {
	log.Printf("sendorder handler executed\nid: %d,\nprice: %d,\nUnits: %v", in.Userid, in.Price, in.Units)

	err := s.orderProvider.ProvideOrder(ctx, in)

	return &grpc_orders.EmptyOrderResp{}, err
}
