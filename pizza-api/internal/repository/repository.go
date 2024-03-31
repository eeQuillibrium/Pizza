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
	//from grpc
	StoreOrder(
		ctx context.Context,
		price int,
		unitNums string,
		amount string,
		state string,
		userId int,
		orderId int,
	) error 
	CancelOrder(
		ctx context.Context,
		userId int,
		orderId int,
	) error
}

type APIProvider interface {
	GetCurrentOrders(
		ctx context.Context,
		userId int,
	) ([]map[string]string, error)
	GetOrdersHistory(
		ctx context.Context,
		userId int,
	) ([]*models.Order, error)
	CancelOrder(
		ctx context.Context,
		orderId int,
		userId int,
	) error
	CreateOrder(
		ctx context.Context,
		price int,
		unit_nums string,
		amount string,
		state string,
		userId int,
	) error 
	StoreReview(
		ctx context.Context,
		userId int,
		reviewText string,
	) error
	StoreUser(
		ctx context.Context,
		address string, 
		email string,
		phone string,
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
