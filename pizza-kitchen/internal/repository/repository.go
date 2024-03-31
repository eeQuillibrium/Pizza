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
	CancelOrder(
		ctx context.Context,
		orderId int,
	) error
}

type Kitchen interface {
	GetCurrentOrders(
		ctx context.Context,
	) []map[string]string
	CancelOrder(
		ctx context.Context,
		orderId int,
	) error
}

type Repository struct {
	OrderProvider
	Kitchen
}

func New(rClient *redis.Client) *Repository {
	return &Repository{
		OrderProvider: NewOPRepo(rClient),
		Kitchen:       NewKitchenRepo(rClient),
	}
}
