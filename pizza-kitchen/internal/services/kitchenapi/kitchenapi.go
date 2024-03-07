package kitchenapi

import (
	"context"

	nikita_kitchen1 "github.com/eeQuillibrium/protos/gen/go/kitchen"
)

type KitchenAPI struct {
}

func New() *KitchenAPI {
	return &KitchenAPI{}
}
func (k *KitchenAPI) SendMessage(
	ctx context.Context,
	in *nikita_kitchen1.SendOrderReq,
) () {}
