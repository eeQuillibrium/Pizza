package server

import (
	"context"

	grpc_orders "github.com/eeQuillibrium/protos/gen/go/orders"
)

// orderprovider server handler implementation
func (s *serverAPI) SendOrder(
	ctx context.Context,
	in *grpc_orders.SendOrderReq,
) (*grpc_orders.EmptyOrderResp, error) {
	return &grpc_orders.EmptyOrderResp{}, s.orderProvider.ProvideOrder(ctx, in)
}
