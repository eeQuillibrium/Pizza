package service

import (
	"context"

	"github.com/eeQuillibrium/pizza-api/internal/domain/models"
	"github.com/eeQuillibrium/pizza-api/internal/logger"
	"github.com/eeQuillibrium/pizza-api/internal/repository"
	grpc_orders "github.com/eeQuillibrium/protos/gen/go/orders"
)

type OPService struct {
	log  *logger.Logger
	repo repository.OrderProvider
}

func NewOPService(
	log *logger.Logger,
	repo repository.OrderProvider,
) *OPService {

	return &OPService{
		log:  log,
		repo: repo,
	}
}
func (s *OPService) ProvideOrder(
	ctx context.Context,
	in *grpc_orders.SendOrderReq,
) error {

	order := &models.Order{
		UserId:  int(in.Userid),
		Price:   int(in.Price),
		OrderId: int(in.Orderid),
		State:   in.GetState().String(),
	}

	for i := 0; i < len(in.Units); i++ {
		order.Units = append(order.Units,
			models.PieceUnitnum{
				Unitnum: int(in.Units[i].Unitnum),
				Piece:   int(in.Units[i].Piece),
			})
	}

	return s.repo.StoreOrder(ctx, order)
}
