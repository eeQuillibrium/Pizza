package service

import (
	"context"

	grpc_orders "github.com/eeQuillibrium/protos/gen/go/orders"

	"github.com/eeQuillibrium/pizza-delivery/internal/domain/models"
	"github.com/eeQuillibrium/pizza-delivery/internal/repository"
)

type OrderProvider interface {
	StoreOrder(
		ctx context.Context,
		in *grpc_orders.SendOrderReq,
	) error
	CancelOrder(
		ctx context.Context,
		in *grpc_orders.SendOrderReq,
	) error
}

type APIProvider interface {
	GetOrders(
		ctx context.Context,
	) ([]*models.Order, error)
	DeleteOrder(
		ctx context.Context,
		orderId int,
	) error
}

type Service struct {
	OrderProvider
	APIProvider
}

func New(
	repo *repository.Repository,
) *Service {

	return &Service{
		OrderProvider: NewOPService(repo.OrderProvider),
		APIProvider:   NewAPIPService(repo.APIProvider),
	}
}
