package service

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/eeQuillibrium/pizza-api/internal/domain/models"
	"github.com/eeQuillibrium/pizza-api/internal/logger"
	"github.com/eeQuillibrium/pizza-api/internal/repository"
)

const unitSep = ","

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
) (int, error) {
	unitNums, amount := s.createUnitRecords(ctx, order.Units)
	return s.repo.CreateOrder(ctx, order.Price, unitNums, amount, order.State, order.UserId)
}
func (s *APIPService) createUnitRecords(
	ctx context.Context,
	units []models.PieceUnitnum,
) (string, string) {
	unitNums := fmt.Sprintf("%d", units[0].Unitnum)
	for i := 1; i < len(units); i++ {
		unitNums += fmt.Sprintf("%s%d", unitSep, units[i].Unitnum)
	}
	pieceNums := fmt.Sprintf("%d", units[0].Piece)
	for i := 1; i < len(units); i++ {
		pieceNums += fmt.Sprintf("%s%d", unitSep, units[i].Piece)
	}
	return unitNums, pieceNums
}

func (s *APIPService) GetCurrentOrders(
	ctx context.Context,
	userId int,
) ([]*models.Order, error) {
	res := []*models.Order{}

	ordersMap, err := s.repo.GetCurrentOrders(ctx, userId)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(ordersMap); i++ {
		order, err := s.getOrder(ordersMap[i])
		if err != nil {
			return nil, err
		}
		res = append(res, order)
	}
	return res, nil
}
func (s *APIPService) getOrder(orderMap map[string]string) (*models.Order, error) {
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
func (s *APIPService) translateUnits(
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

func (s *APIPService) GetOrdersHistory(
	ctx context.Context,
	userId int,
) ([]*models.Order, error) {
	notOrdered, err := s.repo.GetOrdersHistory(ctx, userId)
	if err != nil {
		return nil, err
	}
	s.log.SugaredLogger.Info(notOrdered)
	ordered, err := s.GetCurrentOrders(ctx, userId)
	if err != nil {
		return nil, err
	}
	s.log.SugaredLogger.Info(ordered)

	return append(notOrdered, ordered...), nil
}

func (s *APIPService) CancelOrder(
	ctx context.Context,
	order *models.Order,
) error {
	return s.repo.CancelOrder(ctx, order.OrderId, order.UserId)
}

func (s *APIPService) CreateReview(
	ctx context.Context,
	review *models.Review,
) error {
	return s.repo.StoreReview(ctx, review.UserId, review.Text)
}

func (s *APIPService) StoreUser(
	ctx context.Context,
	user *models.User,
) error {
	return s.repo.StoreUser(ctx, user.Address, user.Email, user.Phone)
}
