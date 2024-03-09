package grpcserver

import (
	"context"
	"log"

	"github.com/eeQuillibrium/pizza-api/internal/service"
	nikita_kitchen1 "github.com/eeQuillibrium/protos/gen/go/kitchen"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	nikita_kitchen1.UnimplementedKitchenServer
	service service.OrderProvider
}

func Register(
	server *grpc.Server,
	service service.OrderProvider,
) {
	nikita_kitchen1.RegisterKitchenServer(server, &GRPCServer{service: service})
}

func (s *GRPCServer) SendOrder(
	ctx context.Context,
	req *nikita_kitchen1.SendOrderReq,
) (*nikita_kitchen1.EmptyOrderResp, error) {
	log.Print("handler in pizzapi was performed")
	return &nikita_kitchen1.EmptyOrderResp{}, nil
}
