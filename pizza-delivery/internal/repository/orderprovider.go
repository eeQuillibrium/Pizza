package repository

import (
	"context"
	"fmt"

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
	orderId int,
	userId int,
	price int,
	state string,
	unitNums string,
	amount string,
) error {
	return r.rClient.HSet(
		ctx,
		fmt.Sprintf("order:%d", orderId),
		"orderid", orderId,
		"userid", userId,
		"price", price,
		"state", state,
		"unitnums", unitNums,
		"amount", amount,
	).Err()
}

func (r *OPRepo) DeleteOrder(
	ctx context.Context,
	orderId int,
) error {
	return r.rClient.Del(ctx, fmt.Sprintf("order:%d", orderId)).Err()
}
