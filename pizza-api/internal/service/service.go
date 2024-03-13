package service

import (
	"context"

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

type Service struct {
	OrderProvider
}

func New(
	log *logger.Logger,
	repo *repository.Repository,
) *Service {
	return &Service{
		OrderProvider: NewOPService(log, repo.OrderProvider),
	}
}
