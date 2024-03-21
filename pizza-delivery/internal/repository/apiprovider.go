package repository

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type APIPRepo struct {
	rClient *redis.Client
}

func NewAPIPRepo(
	rClient *redis.Client,
) *APIPRepo {
	return &APIPRepo{
		rClient: rClient,
	}
}
func (r *APIPRepo) DeleteOrder(
	ctx context.Context,
	orderId int,
) error {
	return r.rClient.Del(ctx, fmt.Sprintf("order:%d", orderId)).Err()
}
func (r *APIPRepo) GetOrders(
	ctx context.Context,
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
