package kitchenapi

import (
	"context"
	"log"

	nikita_kitchen1 "github.com/eeQuillibrium/protos/gen/go/kitchen"
	"google.golang.org/grpc"
)

// for service
type KitchenAPI interface {
	SendMessage(
		ctx context.Context,
		in *nikita_kitchen1.SendOrderReq,
	) (out *nikita_kitchen1.EmptyOrderResp)
}

type serverAPI struct {
	KitchenAPI KitchenAPI
	nikita_kitchen1.UnimplementedKitchenServer
}

func Register(
	serv *grpc.Server,
	service KitchenAPI,
) {
	nikita_kitchen1.RegisterKitchenServer(serv, &serverAPI{KitchenAPI: service})
}

func (s *serverAPI) SendMessage(
	ctx context.Context,
	in *nikita_kitchen1.SendOrderReq,
) (*nikita_kitchen1.EmptyOrderResp, error) {
	log.Print("sendmessage handler executed...")
	return &nikita_kitchen1.EmptyOrderResp{}, nil
}
