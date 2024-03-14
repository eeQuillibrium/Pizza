package repository

import (
	"context"
	"fmt"
	"math/rand"

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

	orderkey := fmt.Sprintf("order:%d", 10e8+rand.Intn(9*10e8-1))

	err := r.rClient.HSet(
		ctx,
		orderkey,
		"userid", order.UserId,
		"price", order.Price,
		"len", len(order.Units),
	).Err()
	if err != nil {
		return err
	}

	for i := 0; i < len(order.Units); i++ {
		err := r.rClient.HSet(ctx, orderkey,
			fmt.Sprintf("unitnum%d", i), order.Units[i].Unitnum,
			fmt.Sprintf("piece%d", i), order.Units[i].Piece).Err()
		if err != nil {
			return err
		}
	}

	return nil
}