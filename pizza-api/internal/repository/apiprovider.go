package repository

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/eeQuillibrium/pizza-api/internal/domain/models"
	"github.com/eeQuillibrium/pizza-api/internal/logger"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

const (
	zeroInt = 0
)

type APIPRepo struct {
	log     *logger.Logger
	rClient *redis.Client
	DB      *sqlx.DB
}

func NewAPIPRepo(
	log *logger.Logger,
	DB *sqlx.DB,
	rClient *redis.Client,
) *APIPRepo {
	return &APIPRepo{
		log:     log,
		DB:      DB,
		rClient: rClient,
	}
}

func (r *APIPRepo) GetCurrentOrders(
	ctx context.Context,
	userId int,
) ([]map[string]string, error) {
	var res []map[string]string

	ordersMap := r.rClient.HGetAll(ctx, fmt.Sprintf("user:%d", userId)).Val()

	for orderKey, orderVal := range ordersMap {
		if orderKey == "last" {
			continue
		}
		res = append(res, r.rClient.HGetAll(ctx, orderVal).Val())
	}

	return res, nil
}

func (r *APIPRepo) CreateOrder(
	ctx context.Context,
	price int,
	unitNums string,
	amount string,
	state string,
	userId int,
) (int, error) {

	orderId, err := r.storeOrder(ctx, price, unitNums, amount, state, userId)
	if err != nil {
		return 0, err
	}

	if err = r.rClient.HSet(
		ctx,
		fmt.Sprintf("order:%d", orderId),
		"orderid", orderId,
		"userid", userId,
		"price", price,
		"state", state,
		"unitnums", unitNums,
		"amount", amount,
	).Err(); err != nil {
		return 0, err
	}

	return orderId, r.updateLast(ctx, userId, orderId)
}
func (r *APIPRepo) storeOrder(
	ctx context.Context,
	price int,
	unitNums string,
	amount string,
	state string,
	userId int,
) (int, error) {
	q := fmt.Sprintf("INSERT INTO %s (price, unit_nums, amount, state, user_id) "+
		"VALUES ($1, $2, $3, $4, $5) RETURNING order_id", "orders")

	var orderId int
	row := r.DB.QueryRow(q, price, unitNums, amount, state, userId)
	if err := row.Scan(&orderId); err != nil {
		return zeroInt, err
	}
	return orderId, nil
}
func (r *APIPRepo) updateLast(
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
		Err()
}

func (r *APIPRepo) CancelOrder(
	ctx context.Context,
	orderId int,
	userId int,
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
func (r *APIPRepo) updateOrder(
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

func (r *APIPRepo) GetOrdersHistory(
	ctx context.Context,
	userId int,
) ([]*models.Order, error) {
	queryOrders := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1 AND state <> 'ORDERED'", "orders")

	type OrderDB struct {
		OrderId  int    `db:"order_id"`
		Price    int    `db:"price"`
		UnitNums string `db:"unit_nums"`
		Amount   string `db:"amount"`
		State    string `db:"state"`
		UserId   int    `db:"user_id"`
	}

	orders := []OrderDB{}

	if err := r.DB.Select(&orders, queryOrders, userId); err != nil {
		return nil, err
	}

	res := []*models.Order{}
	for i := 0; i < len(orders); i++ {
		units, err := accessUnits(orders[i].UnitNums, orders[i].Amount)
		if err != nil {
			return nil, err
		}
		order := &models.Order{
			OrderId: orders[i].OrderId,
			UserId:  userId,
			Price:   orders[i].Price,
			State:   "EXECUTED", // *
			Units:   units,
		}
		res = append(res, order)
	}
	
	return res, nil
}

func accessUnits(unitNumsStr string, amountStr string) ([]models.PieceUnitnum, error) {
	unitNums := strings.Split(unitNumsStr, ",")
	amountNums := strings.Split(amountStr, ",")
	if len(unitNums) != len(amountNums) {
		return nil, errors.New("unmapped units and amounts")
	}

	res := []models.PieceUnitnum{}
	for i := 0; i < len(unitNums); i++ {
		unitnum, err := strconv.Atoi(unitNums[i])
		if err != nil {
			return nil, err
		}
		piece, err := strconv.Atoi(amountNums[i])
		if err != nil {
			return nil, err
		}
		res = append(res, models.PieceUnitnum{Unitnum: unitnum, Piece: piece})
	}

	return res, nil
}

func (r *APIPRepo) StoreReview(
	ctx context.Context,
	userId int,
	reviewText string,
) error {
	query := fmt.Sprintf("INSERT INTO %s (user_id, text) VALUES ($1, $2)", "reviews")

	tx := r.DB.MustBegin()
	tx.MustExecContext(ctx, query, userId, reviewText)
	return tx.Commit()
}

func (r *APIPRepo) StoreUser(
	ctx context.Context,
	address string,
	email string,
	phone string,
) error {
	query := fmt.Sprintf("INSERT INTO %s (address, email, phone) VALUES ($1, $2, $3)", "users")
	tx := r.DB.MustBegin()
	tx.MustExecContext(ctx, query, address, email, phone)
	return tx.Commit()
}
