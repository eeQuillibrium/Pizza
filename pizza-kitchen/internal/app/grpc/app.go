package grpcapp

import (
	"github.com/eeQuillibrium/pizza-kitchen/internal/app/grpc/kitchenapi"
	"github.com/eeQuillibrium/pizza-kitchen/internal/app/grpc/kitchendel"
)

type GRPCApp struct {
	KitchenAPI *kitchenapi.Kitchen
	KitchenDEL *kitchendel.Kitchen
}

func New() *GRPCApp {
	return &GRPCApp{
		kitchenapi.New(),
		kitchendel.New(),
	}
}
