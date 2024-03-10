package repository

import (
	"context"

	"github.com/eeQuillibrium/pizza-api/internal/domain/models"
	"github.com/redis/go-redis/v9"
)

// OrderProvider - OP
type OrderProvider interface {
	StoreOrder(
		ctx context.Context,
		order *models.Order,
	) error
}

type Repository struct {
	OrderProvider
}

func New(rClient *redis.Client) *Repository {
	return &Repository{NewOPRepo(rClient)}
}
