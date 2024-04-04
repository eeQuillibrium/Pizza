package grpcserver

import (
	"context"

	"github.com/eeQuillibrium/pizza-api/internal/service"
	grpc_orders "github.com/eeQuillibrium/protos/gen/go/orders"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	grpc_orders.UnimplementedOrderingServer
	service service.OrderProvider
	port int
}

func Register(
	server *grpc.Server,
	service service.OrderProvider,
) {
	grpc_orders.RegisterOrderingServer(server, &GRPCServer{service: service})
}

func (s *GRPCServer) SendOrder(
	ctx context.Context,
	in *grpc_orders.SendOrderReq,
) (*grpc_orders.EmptyOrderResp, error) {
	if in.GetState().String() == "CANCELLED" {
		return &grpc_orders.EmptyOrderResp{}, s.service.CancelOrder(ctx, in)
	}
	return &grpc_orders.EmptyOrderResp{}, s.service.ProvideOrder(ctx, in)
}
