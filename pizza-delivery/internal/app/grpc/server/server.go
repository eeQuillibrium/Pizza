package grpcserver

import (
	"github.com/eeQuillibrium/pizza-delivery/internal/service"
	grpc_orders "github.com/eeQuillibrium/protos/gen/go/orders"
	"google.golang.org/grpc"
)

type serverAPI struct {
	grpc_orders.UnimplementedOrderingServer
	orderProvider service.OrderProvider
}

func Register(server *grpc.Server, orderProvider service.OrderProvider) {
	grpc_orders.RegisterOrderingServer(server, &serverAPI{orderProvider: orderProvider})
}
