package service

import (
	"context"

	"github.com/eeQuillibrium/pizza-api/internal/domain/models"
	"github.com/eeQuillibrium/pizza-api/internal/logger"
	"github.com/eeQuillibrium/pizza-api/internal/repository"
	grpc_orders "github.com/eeQuillibrium/protos/gen/go/orders"
)

// OrderProvider - OP
type OrderProvider interface {
	ProvideOrder(
		ctx context.Context,
		in *grpc_orders.SendOrderReq,
	) error
}
type APIProvider interface {
	GetOrders(
		ctx context.Context,
	) ([]*models.Order, error)
	CreateOrder(
		context.Context,
		*models.Order,
	) error
}

type Service struct {
	OrderProvider
	APIProvider
}

func New(
	log *logger.Logger,
	repo *repository.Repository,
) *Service {
	return &Service{
		OrderProvider: NewOPService(log, repo.OrderProvider),
		APIProvider:   NewAPIPService(log, repo.APIProvider),
	}
}
