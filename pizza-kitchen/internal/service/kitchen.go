package service

import (
	"context"
	"strconv"
	"strings"

	"github.com/eeQuillibrium/pizza-kitchen/internal/domain/models"
	"github.com/eeQuillibrium/pizza-kitchen/internal/repository"
)

type KitchenService struct {
	repo repository.Kitchen
}

func NewKitchenService(
	repo repository.Kitchen,
) *KitchenService {
	return &KitchenService{repo: repo}
}

func (s *KitchenService) CancelOrder(
	ctx context.Context,
	orderId int,
) error {
	return s.repo.CancelOrder(ctx, orderId)
}
func (s *KitchenService) GetCurrentOrders(
	ctx context.Context,
) ([]*models.Order, error) {
	ordersMap := s.repo.GetCurrentOrders(ctx)

	res := []*models.Order{}
	for i := 0; i < len(ordersMap); i++ {
		order, err := s.getOrder(ordersMap[i])
		if err != nil {
			return res, err
		}
		res = append(res, order)
	}

	return res, nil
}
func (s *KitchenService) getOrder(orderMap map[string]string) (*models.Order, error) {
	price, err := strconv.Atoi(orderMap["price"])
	if err != nil {
		return nil, err
	}
	userId, err := strconv.Atoi(orderMap["userid"])
	if err != nil {
		return nil, err
	}
	orderid, err := strconv.Atoi(orderMap["orderid"])
	if err != nil {
		return nil, err
	}
	units, err := s.translateUnits(orderMap["unitnums"], orderMap["amount"])
	if err != nil {
		return nil, err
	}

	return &models.Order{
		OrderId: orderid,
		Price:   price,
		UserId:  userId,
		Units:   units,
		State:   orderMap["state"],
	}, nil
}
func (s *KitchenService) translateUnits(
	unitNumsS string,
	amountS string,
) ([]models.PieceUnitnum, error) {
	unitNums := strings.Split(unitNumsS, unitSep)
	amount := strings.Split(amountS, unitSep)
	units := []models.PieceUnitnum{}
	for i := 0; i < len(unitNums); i++ {
		unitnum, err := strconv.Atoi(unitNums[i])
		if err != nil {
			return nil, err
		}
		amount, err := strconv.Atoi(amount[i])
		if err != nil {
			return nil, err
		}

		units = append(units, models.PieceUnitnum{Unitnum: unitnum, Piece: amount})
	}
	return units, nil
}
