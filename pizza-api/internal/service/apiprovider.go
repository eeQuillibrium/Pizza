package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/eeQuillibrium/pizza-api/internal/domain/models"
	"github.com/eeQuillibrium/pizza-api/internal/logger"
	"github.com/eeQuillibrium/pizza-api/internal/repository"
)

type APIPService struct {
	log  *logger.Logger
	repo repository.APIProvider
}

func NewAPIPService(
	log *logger.Logger,
	repo repository.APIProvider,
) *APIPService {
	return &APIPService{
		log:  log,
		repo: repo,
	}
}
func (s *APIPService) CreateOrder(
	ctx context.Context,
	order *models.Order,
) error {
	return s.repo.StoreOrder(ctx, order)
}

func (s *APIPService) GetCurrentOrders(
	ctx context.Context,
	userId int,
) ([]*models.Order, error) {

	orders := s.repo.GetCurrentOrders(ctx, userId)
	
	res := []*models.Order{}
	for i := 0; i < len(orders); i++ {
		order, err := getOrder(orders[i])
		if err != nil {
			return res, err
		}
		res = append(res, order)
	}

	return res, nil
}
func getOrder(ordermap map[string]string) (*models.Order, error) {
	price, err := strconv.Atoi(ordermap["price"])
	if err != nil {
		return nil, err
	}
	userId, err := strconv.Atoi(ordermap["userid"])
	if err != nil {
		return nil, err
	}
	len, err := strconv.Atoi(ordermap["len"])
	if err != nil {
		return nil, err
	}
	orderid, err := strconv.Atoi(ordermap["orderid"])
	if err != nil {
		return nil, err
	}
	units := []models.PieceUnitnum{}
	for i := 0; i < len; i++ {
		unitnum, err := strconv.Atoi(ordermap[fmt.Sprintf("unitnum%d", i)])
		if err != nil {
			return nil, err
		}
		piece, err := strconv.Atoi(ordermap[fmt.Sprintf("piece%d", i)])
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
		State:   ordermap["state"],
	}, nil
}

func (s *APIPService) DeleteOrder(
	ctx context.Context,
	order *models.Order,
) error {
	return s.repo.DeleteOrder(ctx, order)
}	
