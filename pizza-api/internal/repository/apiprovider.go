package repository

import (
	"context"
	"fmt"

	"github.com/eeQuillibrium/pizza-api/internal/domain/models"
	"github.com/eeQuillibrium/pizza-api/internal/logger"
	"github.com/redis/go-redis/v9"
)

type APIPRepo struct {
	log     *logger.Logger
	rClient *redis.Client
}

func NewAPIPRepo(
	log *logger.Logger,
	rClient *redis.Client,
) *APIPRepo {
	return &APIPRepo{
		log:     log,
		rClient: rClient,
	}
}

func (r *APIPRepo) GetCurrentOrders(
	ctx context.Context,
	userId int,
) []map[string]string {
	var (
		cursor uint64
		match  string
		count  int64
	)

	res := []map[string]string{}

	keys, _ := r.rClient.Scan(ctx, cursor, match, count).Val()

	for _, key := range keys {
		val := r.rClient.HGetAll(ctx, key).Val()
		res = append(res, val)
	}

	return res
}

func (r *APIPRepo) StoreOrder(
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

	if err := r.rClient.HMSet(ctx, fmt.Sprintf("user:%d", order.UserId), orderkey).
		Err(); err != nil {
		return err
	}

	return nil
}

func (r *APIPRepo) DeleteOrder(
	ctx context.Context,
	order *models.Order,
) error {

	return nil
}
