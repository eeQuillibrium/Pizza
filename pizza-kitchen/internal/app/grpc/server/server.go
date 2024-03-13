package server

import (
	"github.com/eeQuillibrium/pizza-kitchen/internal/service"
	grpc_orders "github.com/eeQuillibrium/protos/gen/go/orders"
	"google.golang.org/grpc"
)

// implemented all service func (handlers)
type serverAPI struct {
	grpc_orders.UnimplementedOrderingServer
	orderProvider service.OrderProvider // service for KitchenServer
}

func Register(
	serv *grpc.Server,
	orderProviderService service.OrderProvider,
) {
	grpc_orders.RegisterOrderingServer(serv, &serverAPI{orderProvider: orderProviderService})
}
