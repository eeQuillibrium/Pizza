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
	GetCurrentOrders(
		ctx context.Context,
		userId int,
	) ([]*models.Order, error)
	GetOrdersHistory(
		ctx context.Context,
		userId int,
	) ([]*models.Order, error)
	CreateOrder(
		ctx context.Context,
		order *models.Order,
	) error
	DeleteOrder(
		ctx context.Context,
		order *models.Order,
	) error
	CreateReview(
		ctx context.Context,
		userId int,
		reviewText string,
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
