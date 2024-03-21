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
	err := s.service.ProvideOrder(ctx, in)
	return &grpc_orders.EmptyOrderResp{}, err
}
