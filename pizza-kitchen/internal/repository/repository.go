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
	GetOrders(ctx context.Context,)
}

type KitchenRedisDB interface {
	OrderProvider
	Kitchen
}

type Repository struct {
	KitchenRedisDB
}

func New(client *redis.Client) *Repository {
	return &Repository{
		KitchenRedisDB: NewRedisDB(client),
	}
}
