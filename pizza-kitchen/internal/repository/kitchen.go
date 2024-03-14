package repository

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type KitchenRepo struct {
	rClient *redis.Client
}

func NewKitchenRepo(rClient *redis.Client) *KitchenRepo {
	return &KitchenRepo{
		rClient: rClient,
	}
}

func (r *KitchenRepo) GetOrders(ctx context.Context) []map[string]string {
	var (
		cursor uint64
		match  string
		count  int64
	)
	res := []map[string]string{}

	scanRes := r.rClient.Scan(ctx, cursor, match, count)
	keys, _ := scanRes.Val()

	for _, key := range keys {
		val := r.rClient.HGetAll(ctx, key).Val()
		res = append(res, val)
	}

	return res
}
