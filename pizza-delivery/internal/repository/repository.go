package repository

import (
	"context"

	"github.com/eeQuillibrium/pizza-delivery/internal/domain/models"
	"github.com/redis/go-redis/v9"
)

type OrderProvider interface {
	StoreOrder(
		ctx context.Context,
		order *models.Order,
	) error

}
type APIProvider interface {
	GetOrders(
		ctx context.Context,
	) []map[string]string
}

type Repository struct {
	OrderProvider
	APIProvider
}

func New(
	rClient *redis.Client,
) *Repository {
	return &Repository{
		OrderProvider: NewOPRepo(rClient),
		APIProvider:   NewAPIPRepo(rClient),
	}
}
