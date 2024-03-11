package server

import (
	"github.com/eeQuillibrium/pizza-kitchen/internal/service"
	nikita_kitchen1 "github.com/eeQuillibrium/protos/gen/go/kitchen"
	"google.golang.org/grpc"
)

// implemented all service func (handlers)
type serverAPI struct {
	nikita_kitchen1.UnimplementedKitchenServer
	orderProvider service.OrderProvider // service for KitchenServer
}

func Register(
	serv *grpc.Server,
	orderProviderService service.OrderProvider,
) {
	nikita_kitchen1.RegisterKitchenServer(serv, &serverAPI{orderProvider: orderProviderService})
}
