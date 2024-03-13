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
	r.log.SugaredLogger.Info("try to store order in redis", order.UserId, order.Price, order.Units)

	orderkey := fmt.Sprintf("order:%d", 10e8+rand.Intn(9*10e8-1))
	err := r.rClient.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		r.log.SugaredLogger.Fatal(err)
	}

	err = r.rClient.HSet(
		ctx,
		orderkey,
		"userid", order.UserId,
		"price", order.Price,
	).Err()
	if err != nil {
		r.log.SugaredLogger.Info("unsuccessful order storing")
		return err
	}

	for i := 0; i < len(order.Units); i++ {
		err := r.rClient.HSet(ctx, orderkey,
			fmt.Sprintf("unitnum%d", i), order.Units[i].Unitnum,
			fmt.Sprintf("piece%d", i), order.Units[i].Piece).Err()
		if err != nil {
			r.log.SugaredLogger.Info("unsuccessful order storing")
			return err
		}
	}

	r.log.SugaredLogger.Info("successful order storing!")

	return nil
}
