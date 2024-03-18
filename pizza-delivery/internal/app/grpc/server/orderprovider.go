package grpcserver

import (
	"context"

	grpc_orders "github.com/eeQuillibrium/protos/gen/go/orders"
)

func (s *serverAPI) SendOrder(
	ctx context.Context,
	in *grpc_orders.SendOrderReq,
) (*grpc_orders.EmptyOrderResp, error) {
	return s.orderProvider.StoreOrder(ctx, in)
}
