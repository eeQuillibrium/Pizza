package repository

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/eeQuillibrium/pizza-kitchen/internal/domain/models"
	"github.com/redis/go-redis/v9"
)

type RedisDB struct {
	rClient *redis.Client
}

func NewRedisDB(rClient *redis.Client) *RedisDB {
	return &RedisDB{rClient: rClient}
}

func (r *RedisDB) StoreOrder(
	ctx context.Context,
	order *models.Order,
) error {
	log.Print("try to store order in redis")

	orderkey := fmt.Sprintf("order:%d", 10e8+rand.Intn(9*10e8-1))

	err := r.rClient.HSet(
		ctx,
		orderkey,
		"userid", order.UserId,
		"price", order.Price,
	).Err()
	if err != nil {
		log.Print("unsuccessful order storing")
		return err
	}

	for i := 0; i < len(order.Units); i++ {
		err := r.rClient.HSet(ctx, orderkey,
			fmt.Sprintf("unitnum%d", i), order.Units[i].Unitnum,
			fmt.Sprintf("piece%d", i), order.Units[i].Piece).Err()
		if err != nil {
			log.Print("unsuccessful order storing")
			return err
		}
	}

	log.Print("successful order storing!")

	return nil
}

func (r *RedisDB) GetOrders(ctx context.Context) ([]*models.Order, error) {
	return nil, nil
}
