package service

import (
	"context"

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
	return s.repo.GetCurrentOrders(ctx, userId)
}

func (s *APIPService) GetOrdersHistory(
	ctx context.Context,
	userId int,
) ([]*models.Order, error) {
	return s.repo.GetOrdersHistory(ctx, userId)
}

func (s *APIPService) DeleteOrder(
	ctx context.Context,
	order *models.Order,
) error {
	return s.repo.DeleteOrder(ctx, order)
}
func (s *APIPService) CreateReview(
	ctx context.Context,
	userId int,
	reviewText string,
) error {
	return s.repo.StoreReview(ctx, userId, reviewText)
}
