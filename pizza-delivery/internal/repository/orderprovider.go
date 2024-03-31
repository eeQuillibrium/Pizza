package repository

import (
	"context"
	"fmt"

	"github.com/eeQuillibrium/pizza-delivery/internal/domain/models"
	"github.com/redis/go-redis/v9"
)

type OPRepo struct {
	rClient *redis.Client
}

func NewOPRepo(
	rClient *redis.Client,
) *OPRepo {
	return &OPRepo{
		rClient: rClient,
	}
}
func (r *OPRepo) StoreOrder(
	ctx context.Context,
	order *models.Order,
) error {
	orderKey := fmt.Sprintf("order:%d", order.OrderId)
	if err := r.rClient.HMSet(
		ctx,
		orderKey,
		"orderid", order.OrderId,
		"userid", order.UserId,
		"price", order.Price,
		"state", order.State,
		"len", len(order.Units),
	).Err(); err != nil {
		return err
	}

	for i := 0; i < len(order.Units); i++ {
		if err := r.rClient.HMSet(ctx, orderKey,
			fmt.Sprintf("unitnum%d", i), order.Units[i].Unitnum,
			fmt.Sprintf("piece%d", i), order.Units[i].Piece).
			Err(); err != nil {
			return err
		}
	}

	return nil
}

func (r *OPRepo) DeleteOrder(
	ctx context.Context,
	orderId int,
) error {
	return r.rClient.Del(ctx, fmt.Sprintf("order:%d", orderId)).Err()
}
