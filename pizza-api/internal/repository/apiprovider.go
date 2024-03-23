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
) ([]*models.Order, error) {
	res := []*models.Order{}

	ordersMap := r.rClient.HGetAll(ctx, fmt.Sprintf("user:%d", userId)).Val()

	for _, orderKey := range ordersMap {
		order, err := getOrder(r.rClient.HGetAll(ctx, orderKey).Val())
		if err != nil {
			return res, err
		}
		res = append(res, order)
	}

	return res, nil
}
func getOrder(orderMap map[string]string) (*models.Order, error) {
	price, err := strconv.Atoi(orderMap["price"])
	if err != nil {
		return nil, err
	}
	userId, err := strconv.Atoi(orderMap["userid"])
	if err != nil {
		return nil, err
	}
	len, err := strconv.Atoi(orderMap["len"])
	if err != nil {
		return nil, err
	}
	orderid, err := strconv.Atoi(orderMap["orderid"])
	if err != nil {
		return nil, err
	}
	units := []models.PieceUnitnum{}
	for i := 0; i < len; i++ {
		unitnum, err := strconv.Atoi(orderMap[fmt.Sprintf("unitnum%d", i)])
		if err != nil {
			return nil, err
		}
		piece, err := strconv.Atoi(orderMap[fmt.Sprintf("piece%d", i)])
		if err != nil {
			return nil, err
		}
		units = append(units, models.PieceUnitnum{
			Unitnum: unitnum,
			Piece:   piece,
		})
	}
	return &models.Order{
		OrderId: orderid,
		Price:   price,
		UserId:  userId,
		Units:   units,
		State:   orderMap["state"],
	}, nil
}

func (r *APIPRepo) StoreOrder(
	ctx context.Context,
	order *models.Order,
) error {
	var err error
	orderKey := fmt.Sprintf("order:%d", order.OrderId)

	if err = r.storeOrder(ctx, order, orderKey); err != nil {
		return err
	}

	userKey := fmt.Sprintf("user:%d", order.UserId)

	var last int = -1

	lastStr := r.rClient.HGetAll(ctx, userKey).Val()["last"]
	if lastStr != "" {
		last, err = strconv.Atoi(lastStr)
		if err != nil {
			return err
		}
	}

	if err := r.rClient.HMSet(ctx, userKey,
		fmt.Sprintf("order:%d", last+1), orderKey,
		"last", last+1).
		Err(); err != nil {
		return err
	}

	return nil
}
func (r *APIPRepo) storeOrder(
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

func (r *APIPRepo) DeleteOrder(
	ctx context.Context,
	order *models.Order,
) error {

	return nil
}

func (r *APIPRepo) GetOrdersHistory(
	ctx context.Context,
	userId int,
) ([]*models.Order, error) {
	queryOrders := fmt.Sprintf("SELECT * FROM %s WHERE user_id = '$1'", "userOrder")

	type OrderDB struct {
		OrderId  int    `db:"order_id"`
		Price    int    `db:"price"`
		UnitNums string `db:"unit_nums"`
		Amount   string `db:"amount"`
	}
	orders := []OrderDB{}

	err := r.DB.GetContext(ctx, &orders, queryOrders, userId)
	if err != nil {
		return nil, nil
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
	query := fmt.Sprintf("INSERT INTO %s (user_id, review) VALUES ($1, $2)", "reviews")

	tx := r.DB.MustBegin()

	tx.MustExecContext(ctx, query, userId, reviewText)

	return tx.Commit()
}
