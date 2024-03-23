package repository

import (
	"context"

	"github.com/eeQuillibrium/pizza-api/internal/domain/models"
	"github.com/eeQuillibrium/pizza-api/internal/logger"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

// OrderProvider - OP
type OrderProvider interface {
	StoreOrder(
		ctx context.Context,
		order *models.Order,
	) error //from grpc
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
	DeleteOrder(
		ctx context.Context,
		order *models.Order,
	) error
	StoreOrder(
		ctx context.Context,
		order *models.Order,
	) error //order from client
	StoreReview(
		ctx context.Context,
		userId int,
		reviewText string,
	) error
}

type Repository struct {
	OrderProvider
	APIProvider
}

func New(
	log *logger.Logger,
	db *sqlx.DB,
	rClient *redis.Client,
) *Repository {
	return &Repository{
		OrderProvider: NewOPRepo(log, db, rClient),
		APIProvider:   NewAPIPRepo(log, db, rClient),
	}
}
