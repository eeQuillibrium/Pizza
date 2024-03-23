package repository

import (
	"context"
	"fmt"
	"strconv"

	"github.com/eeQuillibrium/pizza-api/internal/domain/models"
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
	order *models.Order,
) error {
	if order.State == "EXECUTED" {
		queryUser := fmt.Sprintf("INSERT INTO %s (user_id, order_id) VALUES ($1, $2)", "userOrder")
		queryOrder := fmt.Sprintf("INSERT INTO %s (order_id, price, unit_nums, amount) VALUES ($1 $2 $3 $4)", "orders")
		unitNums, amount := createRecords(ctx, order.Units)
		//unit_num is a sequence of unitnums wit seq=","

		tx := r.DB.MustBegin()

		tx.MustExecContext(ctx, queryUser, order.UserId, order.OrderId)
		tx.MustExecContext(ctx, queryOrder, order.OrderId, order.Price, unitNums, amount)

		return tx.Commit()
	}
	var err error
	orderKey := fmt.Sprintf("order:%d", order.OrderId)

	if err = r.storeOrder(ctx, order, orderKey); err != nil {
		return err
	}

	userKey := fmt.Sprintf("user:%d", order.UserId)

	var last int = -1

	lastStr := r.rClient.HGetAll(ctx, userKey).Val()["last"]
	if lastStr != "" {
		last, err = strconv.Atoi(lastStr) // *
		if err != nil {
			return err
		}
	}

	err = r.rClient.HMSet(ctx, userKey,
		fmt.Sprintf("order:%d", last+1), orderKey,
		"last", last+1).Err()

	return err
}
func (r *OPRepo) storeOrder(
	ctx context.Context,
	order *models.Order,
	orderKey string,
) error {
	if err := r.rClient.HMSet(
		ctx,
		orderKey,
		"orderid", order.OrderId,
		"userid", order.UserId,
		"price", order.Price,
		"state", order.State,
		"len", len(order.Units),
	).Err(); err != nil {
		return err
	}

	for i := 0; i < len(order.Units); i++ {
		if err := r.rClient.HMSet(ctx, orderKey,
			fmt.Sprintf("unitnum%d", i), order.Units[i].Unitnum,
			fmt.Sprintf("piece%d", i), order.Units[i].Piece).
			Err(); err != nil {
			return err
		}
	}

	return nil
}
func createRecords(
	ctx context.Context,
	units []models.PieceUnitnum,
) (string, string) {
	unitNums := ""
	for i := 0; i < len(units)-1; i++ {
		unitNums += fmt.Sprintf("%d,", units[i].Unitnum)
	}
	pieceNums := ""
	for i := 0; i < len(units)-1; i++ {
		pieceNums += fmt.Sprintf("%d,", units[i].Piece)
	}
	return unitNums, pieceNums
}
