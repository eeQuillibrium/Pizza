package service

import (
	"context"

	"github.com/eeQuillibrium/pizza-api/internal/domain/models"
	"github.com/eeQuillibrium/pizza-api/internal/repository"
	nikita_kitchen1 "github.com/eeQuillibrium/protos/gen/go/kitchen"
)

type OPService struct {
	repo repository.OrderProvider
}

func NewOPService(
	repo repository.OrderProvider,
) *OPService {
	return &OPService{repo: repo}
}
func (s *OPService) ProvideOrder(
	ctx context.Context,
	in *nikita_kitchen1.SendOrderReq,
) error {

	order := &models.Order{
		UserId: in.Userid,
		Price:  in.Price,
	}

	for i := 0; i < len(in.Units); i++ {
		order.Units = append(order.Units,
			models.PieceUnitnum{
				Unitnum: in.Units[i].Unitnum,
				Piece:   in.Units[i].Piece,
			})
	}

	return s.repo.StoreOrder(ctx, order)
}
