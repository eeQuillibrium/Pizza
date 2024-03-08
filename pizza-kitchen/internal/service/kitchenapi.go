package service

import (
	"context"

	"github.com/eeQuillibrium/pizza-kitchen/internal/domain/models"
	"github.com/eeQuillibrium/pizza-kitchen/internal/repository"

	nikita_kitchen1 "github.com/eeQuillibrium/protos/gen/go/kitchen"
)

type KitchenAPIService struct {
	repo repository.KitchenRedisDB
}

func NewKitchenAPIService(
	repo repository.KitchenRedisDB,
) *KitchenAPIService {
	return &KitchenAPIService{repo: repo}
}

func (s *KitchenAPIService) SendMessage(
	ctx context.Context,
	in *nikita_kitchen1.SendOrderReq,
) error {
	order := &models.Order{
		UserId: in.Userid,
		Price:  in.Price,
	}

	for i := 0; i < len(in.Units); i++ {
		order.Units = append(order.Units,
			&models.PieceUnitnum{
				Unitnum: in.Units[i].Unitnum,
				Piece:   in.Units[i].Piece,
			})
	}

	err := s.repo.CreateOrder(ctx, order)
	
	if err != nil {
		return err
	}
	

	return nil
}
