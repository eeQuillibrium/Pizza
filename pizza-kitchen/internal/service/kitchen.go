package service

import "github.com/eeQuillibrium/pizza-kitchen/internal/repository"

type KitchenService struct {
	repo repository.KitchenRedisDB
}

func NewKitchenService(
	repo repository.KitchenRedisDB,
) *KitchenService {
	return &KitchenService{repo: repo}
}
func (s *KitchenService) GetOrders() {

}
