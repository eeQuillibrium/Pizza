package repository

import (
	"context"

	"github.com/eeQuillibrium/pizza-api/internal/domain/models"
	"github.com/eeQuillibrium/pizza-api/internal/logger"
	"github.com/redis/go-redis/v9"
)

// OrderProvider - OP
type OrderProvider interface {
	StoreOrder(
		ctx context.Context,
		order *models.Order,
	) error
}

type APIProvider interface {
	GetCurrentOrders(
		ctx context.Context,
		userId int,
	) []map[string]string
	StoreOrder(
		ctx context.Context,
		order *models.Order,
	) error
	DeleteOrder(
		ctx context.Context,
		order *models.Order,
	) error
}

type Repository struct {
	OrderProvider
	APIProvider
}

func New(
	log *logger.Logger,
	rClient *redis.Client,
) *Repository {
	return &Repository{
		OrderProvider: NewOPRepo(log, rClient),
		APIProvider:   NewAPIPRepo(log, rClient),
	}
}
