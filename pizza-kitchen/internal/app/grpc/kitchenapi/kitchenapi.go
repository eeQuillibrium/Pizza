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
	)
}

// implemented all service func (handlers)
type serverAPI struct {
	nikita_kitchen1.UnimplementedKitchenServer
	kitchenAPI KitchenAPI // service
}

func Register(
	serv *grpc.Server,
	service KitchenAPI,
) {
	nikita_kitchen1.RegisterKitchenServer(serv, &serverAPI{kitchenAPI: service})
}
func (s *serverAPI) SendOrder(
	ctx context.Context,
	r *nikita_kitchen1.SendOrderReq,
) (*nikita_kitchen1.EmptyOrderResp, error) {
	log.Printf("sendmessage handler executed... id: %d, price: %d", r.Userid, r.Price)
	return &nikita_kitchen1.EmptyOrderResp{}, nil
}
