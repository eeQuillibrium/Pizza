package grpcserver

import (
	"context"

	grpc_orders "github.com/eeQuillibrium/protos/gen/go/orders"
)

func (s *serverAPI) SendOrder(
	ctx context.Context,
	in *grpc_orders.SendOrderReq,
) (*grpc_orders.EmptyOrderResp, error) {
	resp := &grpc_orders.EmptyOrderResp{}
	if in.GetState().String() == "CANCELLED" {
		return resp, s.orderProvider.CancelOrder(ctx, int(in.GetOrderid()))
	}
	return resp, s.orderProvider.StoreOrder(ctx, in)
}
