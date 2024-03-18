package service

import (
	"context"

	"github.com/eeQuillibrium/pizza-delivery/internal/domain/models"
	"github.com/eeQuillibrium/pizza-delivery/internal/repository"
	grpc_orders "github.com/eeQuillibrium/protos/gen/go/orders"
)

type OPService struct {
	repo repository.OrderProvider
}

func NewOPService(
	repo repository.OrderProvider,
) *OPService {
	return &OPService{
		repo: repo,
	}
}

func (s *OPService) StoreOrder(
	ctx context.Context,
	in *grpc_orders.SendOrderReq,
) (*grpc_orders.EmptyOrderResp, error) {
	order := &models.Order{
		OrderId: int(in.Orderid),
		UserId:  int(in.Userid),
		Price:   int(in.Price),
		State:   in.GetState().String(),
	}

	for i := 0; i < len(in.Units); i++ {
		order.Units = append(order.Units,
			models.PieceUnitnum{
				Unitnum: int(in.Units[i].Unitnum),
				Piece:   int(in.Units[i].Piece),
			})
	}

	return &grpc_orders.EmptyOrderResp{}, s.repo.StoreOrder(ctx, order)

}
