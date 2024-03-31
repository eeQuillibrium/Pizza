package repository

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type OrderProvider interface {
	StoreOrder(
		ctx context.Context,
		orderId int,
		userId int,
		price int,
		state string,
		unitNums string,
		amount string,
	) error
	DeleteOrder(
		ctx context.Context,
		orderId int,
	) error
}
type APIProvider interface {
	GetCurrentOrders(
		ctx context.Context,
	) []map[string]string
	DeleteOrder(
		ctx context.Context,
		orderId int,
	) error
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
