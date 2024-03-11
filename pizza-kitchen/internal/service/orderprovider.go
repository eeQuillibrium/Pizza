package service

import (
	"context"
	"log"

	"github.com/eeQuillibrium/pizza-kitchen/internal/domain/models"
	"github.com/eeQuillibrium/pizza-kitchen/internal/repository"

	nikita_kitchen1 "github.com/eeQuillibrium/protos/gen/go/kitchen"
)

type OPService struct {
	repo repository.KitchenRedisDB
}

func NewOPService(
	repo repository.KitchenRedisDB,
) *OPService {
	return &OPService{repo: repo}
}

func (s *OPService) ProvideOrder(
	ctx context.Context,
	in *nikita_kitchen1.SendOrderReq,
) error {
	order := &models.Order{
		UserId: int(in.Userid),
		Price:  int(in.Price),
	}
	log.Print("try to store order", in.Units)
	for i := 0; i < len(in.Units); i++ {
		order.Units = append(order.Units,
			models.PieceUnitnum{
				Unitnum: int(in.Units[i].Unitnum),
				Piece:   int(in.Units[i].Piece),
			})
	}

	err := s.repo.StoreOrder(ctx, order)

	if err != nil {
		return err
	}

	return nil
}
