package repository

import (
	"context"

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

func (r *APIPRepo) GetOrders(
	ctx context.Context,
) []map[string]string {
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
