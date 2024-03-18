package repository

import (
	"context"
	"fmt"

	"github.com/eeQuillibrium/pizza-api/internal/domain/models"
	"github.com/eeQuillibrium/pizza-api/internal/logger"
	"github.com/redis/go-redis/v9"
)

type OPRepo struct {
	log     *logger.Logger
	rClient *redis.Client
}

func NewOPRepo(
	log *logger.Logger,
	rClient *redis.Client,
) *OPRepo {
	return &OPRepo{
		log:     log,
		rClient: rClient,
	}
}

func (r *OPRepo) StoreOrder(
	ctx context.Context,
	order *models.Order,
) error {
	orderkey := fmt.Sprintf("order:%d", order.OrderId)
	if err := r.rClient.HMSet(
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
		if err := r.rClient.HMSet(ctx, orderkey,
			fmt.Sprintf("unitnum%d", i), order.Units[i].Unitnum,
			fmt.Sprintf("piece%d", i), order.Units[i].Piece).
			Err(); err != nil {
			return err
		}
	}

	return nil
}
