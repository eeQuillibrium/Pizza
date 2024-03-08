package server

import (
	"context"
	"log"

	nikita_kitchen1 "github.com/eeQuillibrium/protos/gen/go/kitchen"
)

func (s *serverAPI) SendOrder(
	ctx context.Context,
	r *nikita_kitchen1.SendOrderReq,
) (*nikita_kitchen1.EmptyOrderResp, error) {
	log.Printf("sendmessage handler executed... id: %d, price: %d, Units: %v", r.Userid, r.Price, r.Units)

	s.kitchenAPI.SendMessage(ctx, r)
	
	return &nikita_kitchen1.EmptyOrderResp{}, nil
}
