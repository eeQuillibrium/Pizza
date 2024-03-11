package service

import (
	"context"

	"github.com/eeQuillibrium/pizza-kitchen/internal/repository"
	nikita_kitchen1 "github.com/eeQuillibrium/protos/gen/go/kitchen"
)

type OrderProvider interface {
	ProvideOrder(
		ctx context.Context,
		in *nikita_kitchen1.SendOrderReq,
	) error
}
type Kitchen interface {
	GetOrders()
}

type KitchenRedisService interface {
	OrderProvider
	Kitchen
}

type Service struct {
	OrderProvider
	Kitchen
}

func New(repo *repository.Repository) *Service {
	return &Service{
		OrderProvider: NewOPService(repo.KitchenRedisDB),
		Kitchen:    NewKitchenService(repo.KitchenRedisDB),
	}
}
