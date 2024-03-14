package service

import (
	"context"

	"github.com/eeQuillibrium/pizza-kitchen/internal/domain/models"
	"github.com/eeQuillibrium/pizza-kitchen/internal/repository"
	grpc_orders "github.com/eeQuillibrium/protos/gen/go/orders"
)

type OrderProvider interface {
	ProvideOrder(
		ctx context.Context,
		in *grpc_orders.SendOrderReq,
	) error
}
type Kitchen interface {
	GetOrders(ctx context.Context) ([]*models.Order, error)
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
		OrderProvider: NewOPService(repo.OrderProvider),
		Kitchen:       NewKitchenService(repo.Kitchen),
	}
}
