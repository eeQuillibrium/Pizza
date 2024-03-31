package repository

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/eeQuillibrium/pizza-api/internal/logger"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type OPRepo struct {
	log     *logger.Logger
	DB      *sqlx.DB
	rClient *redis.Client
}

func NewOPRepo(
	log *logger.Logger,
	db *sqlx.DB,
	rClient *redis.Client,
) *OPRepo {
	return &OPRepo{
		log:     log,
		DB:      db,
		rClient: rClient,
	}
}

func (r *OPRepo) StoreOrder(
	ctx context.Context,
	price int,
	unitNums string,
	amount string,
	state string,
	userId int,
	orderId int,
) error {
	if state == "EXECUTED" {
		query := fmt.Sprintf("UPDATE orders SET state = '%s' WHERE order_id = $1", state)
		_, err := r.DB.ExecContext(ctx, query, orderId)
		return err
	}

	if err := r.rClient.HSet(
		ctx,
		fmt.Sprintf("order:%d", orderId),
		"orderid", orderId,
		"userid", userId,
		"price", price,
		"state", state,
		"unitnums", unitNums,
		"amount", amount,
	).Err(); err != nil {
		return err
	}

	return r.updateLast(ctx, userId , orderId)
}
func (r *OPRepo) updateLast(
	ctx context.Context,
	userId int,
	orderId int,
) error {
	userKey := fmt.Sprintf("user:%d", userId)
	var err error
	var last int = -1
	if is := r.rClient.Exists(ctx, userKey).Val(); is == 1 {
		last, err = strconv.Atoi(r.rClient.HGet(ctx, userKey, "last").Val())
		if err != nil {
			return err
		}
	}

	return r.rClient.HSet(ctx, userKey,
		fmt.Sprintf("order:%d", last+1), fmt.Sprintf("order:%d", orderId),
		"last", last+1).
		Err();
}

func (r *OPRepo) CancelOrder(
	ctx context.Context,
	userId int,
	orderId int,
) error {
	if err := r.updateOrder(ctx, orderId); err != nil {
		return err
	}
	userKey := fmt.Sprintf("user:%d", userId)
	orderKey := fmt.Sprintf("order:%d", orderId)

	if err := r.rClient.HDel(ctx, userKey, orderKey).
		Err(); err != nil {
		return err
	}
	if err := r.rClient.Del(ctx, orderKey).
		Err(); err != nil {
		return err
	}

	lastVal := r.rClient.HGet(ctx, userKey, "last").Val()
	if lastVal == "0" {
		return r.rClient.Del(ctx, userKey).Err()
	}

	last, err := strconv.Atoi(lastVal)
	if err != nil {
		return err
	}

	return r.rClient.HSet(ctx, userKey, "last", last-1).Err()
}
func (r *OPRepo) updateOrder(
	ctx context.Context,
	orderId int,
) error {
	if isExist := r.rClient.Exists(ctx, fmt.Sprintf("order:%d", orderId)).
		Val(); isExist == 0 {
		return errors.New("there's no order with this id")
	}

	query := fmt.Sprintf("UPDATE %s SET state = 'CANCELLED'"+
		" WHERE order_id = $1", "orders")

	tx := r.DB.MustBegin()
	tx.MustExecContext(ctx, query, orderId)
	return tx.Commit()
}