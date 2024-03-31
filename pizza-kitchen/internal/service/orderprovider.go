package service

import (
	"context"
	"fmt"

	"github.com/eeQuillibrium/pizza-kitchen/internal/repository"

	grpc_orders "github.com/eeQuillibrium/protos/gen/go/orders"
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
	in *grpc_orders.SendOrderReq,
) error {
	unitNums, amount := s.createUnitRecords(ctx, in.GetUnits())

	return s.repo.StoreOrder(
		ctx,
		int(in.GetOrderid()),
		int(in.GetUserid()),
		int(in.GetPrice()),
		in.GetState().String(),
		unitNums,
		amount,
	)
}
func (s *OPService) createUnitRecords(
	ctx context.Context,
	units []*grpc_orders.SendOrderReq_PieceUnitnum,
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

func (s *OPService) CancelOrder(
	ctx context.Context,
	orderId int,
) error {
	return s.repo.CancelOrder(
		ctx,
		orderId,
	)
}
