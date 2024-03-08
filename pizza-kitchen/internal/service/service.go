package service

import (
	"context"

	"github.com/eeQuillibrium/pizza-kitchen/internal/repository"
	nikita_kitchen1 "github.com/eeQuillibrium/protos/gen/go/kitchen"
)

type KitchenAPI interface {
	SendMessage(
		ctx context.Context,
		in *nikita_kitchen1.SendOrderReq,
	) error
}
type Kitchen interface {
	GetOrders()
}

type KitchenRedisService interface {
	KitchenAPI
	Kitchen
}

type Service struct {
	KitchenAPI
	Kitchen
}

func New(repo *repository.Repository) *Service {
	return &Service{
		KitchenAPI: NewKitchenAPIService(repo.KitchenRedisDB),
		Kitchen:    NewKitchenService(repo.KitchenRedisDB),
	}
}
