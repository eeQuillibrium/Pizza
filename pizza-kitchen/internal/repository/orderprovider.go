package repository

import (
	"context"
	"fmt"

	"github.com/eeQuillibrium/pizza-kitchen/internal/domain/models"
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
	orderkey := fmt.Sprintf("order:%d", order.OrderId)
	if err := r.rClient.HSet(
		ctx,
		orderkey,
		"orderid", order.OrderId,
		"userid", order.UserId,
		"price", order.Price,
		"state", order.State,
		"len", len(order.Units),
	).Err(); err != nil {
		return err
	}

	for i := 0; i < len(order.Units); i++ {
		if err := r.rClient.HSet(ctx, orderkey,
			fmt.Sprintf("unitnum%d", i), order.Units[i].Unitnum,
			fmt.Sprintf("piece%d", i), order.Units[i].Piece).
			Err(); err != nil {
			return err
		}
	}
	if err := r.rClient.HMSet(ctx, fmt.Sprintf("user:%d", order.UserId), orderkey).
		Err(); err != nil {
		return err
	}

	return nil
}
