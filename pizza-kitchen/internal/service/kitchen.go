package service

import (
	"context"

	"github.com/eeQuillibrium/pizza-kitchen/internal/domain/models"
	"github.com/eeQuillibrium/pizza-kitchen/internal/repository"
)

type KitchenService struct {
	repo repository.KitchenRedisDB
}

func NewKitchenService(
	repo repository.KitchenRedisDB,
) *KitchenService {
	return &KitchenService{repo: repo}
}
func (s *KitchenService) GetOrders(
	ctx context.Context,
) ([]*models.Order, error) {
	return s.repo.GetOrders(ctx)
}
