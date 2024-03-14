package repository

import (
	"context"

	"github.com/eeQuillibrium/pizza-kitchen/internal/domain/models"
	"github.com/redis/go-redis/v9"
)

type OrderProvider interface {
	StoreOrder(
		ctx context.Context,
		order *models.Order,
	) error
}

type Kitchen interface {
	GetOrders(ctx context.Context) []map[string]string
}

type Repository struct {
	OrderProvider
	Kitchen
}

func New(rClient *redis.Client) *Repository {
	return &Repository{
		OrderProvider: NewOPRepo(rClient),
		Kitchen: NewKitchenRepo(rClient),
	}
}
