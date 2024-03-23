package repository

import (
	"context"
	"fmt"
	"strconv"

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
	var err error
	orderKey := fmt.Sprintf("order:%d", order.OrderId)

	if err = r.storeOrder(ctx, order, orderKey); err != nil {
		return err
	}

	userKey := fmt.Sprintf("user:%d", order.UserId)

	var last int = -1

	lastStr := r.rClient.HGetAll(ctx, userKey).Val()["last"]
	if lastStr != "" {
		last, err = strconv.Atoi(lastStr)
		if err != nil {
			return err
		}
	}

	if err := r.rClient.HMSet(
		ctx, userKey,
		fmt.Sprintf("order:%d", last+1), orderKey,
		"last", last+1).
		Err(); err != nil {
		return err
	}

	return nil
}
func (r *OPRepo) storeOrder(
	ctx context.Context,
	order *models.Order,
	orderKey string,
) error {
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
