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
	Client *redis.Client
}

func NewRedisDB(client *redis.Client) *RedisDB {
	return &RedisDB{Client: client}
}

func (r *RedisDB) CreateOrder(
	ctx context.Context,
	order *models.Order,
) error {
	log.Print("try to create order in redis...")

	orderkey := fmt.Sprintf("order:%d", 10e8+rand.Intn(9*10e8-1))
	err := r.Client.HSet(
		ctx,
		orderkey,
		"userid", order.UserId,
		"price", order.Price,
	).Err()
	if err != nil {
		log.Print("unsuccessful order creation")
		return err
	}

	for i := 0; i < len(order.Units); i++ {
		err := r.Client.HSet(ctx, orderkey, order.Units[i].Unitnum, order.Units[i].Piece).Err()
		if err != nil {
			log.Print("unsuccessful order creation")
			return err
		}
	}

	log.Print("successful order creation!")

	order_Test(ctx, r, orderkey)

	return nil
}
func (r *RedisDB) GetOrders(ctx context.Context) {

}
func order_Test(
	ctx context.Context,
	repo *RedisDB,
	orderkey string,
) {
	record := repo.Client.HGetAll(ctx, orderkey)
	log.Println(record)
}
