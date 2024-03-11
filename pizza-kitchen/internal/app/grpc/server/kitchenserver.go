package server

import (
	"context"
	"log"

	nikita_kitchen1 "github.com/eeQuillibrium/protos/gen/go/kitchen"
)

// orderprovider server handler implementation
func (s *serverAPI) SendOrder(
	ctx context.Context,
	in *nikita_kitchen1.SendOrderReq,
) (*nikita_kitchen1.EmptyOrderResp, error) {
	log.Printf("sendorder handler executed... id: %d, price: %d, Units: %v", in.Userid, in.Price, in.Units)

	s.orderProvider.ProvideOrder(ctx, in)

	return &nikita_kitchen1.EmptyOrderResp{}, nil
}
