package repository

import (
	"context"
	"fmt"
	"math/rand"

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
	orderkey := generateOrderKey()

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

func generateOrderKey() string {
	return fmt.Sprintf("order:%d", 10e8+rand.Intn(9*10e8-1))
}
