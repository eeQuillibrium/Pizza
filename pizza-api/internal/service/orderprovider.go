package service

import (
	"context"
	"fmt"

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
	unitNums, amounts := s.createUnitRecords(ctx, in.GetUnits())

	return s.repo.StoreOrder(
		ctx, 
		int(in.GetPrice()),
		unitNums,
		amounts,
		in.GetState().String(),
		int(in.GetUserid()),
		int(in.GetOrderid()),
	)
}
func (s *OPService) createUnitRecords(
	ctx context.Context,
	units []*grpc_orders.SendOrderReq_PieceUnitnum,
) (string, string) {
	unitNums := fmt.Sprintf("%d",units[0].Unitnum)
	for i := 1; i < len(units); i++ {
		unitNums += fmt.Sprintf("%s%d", unitSep, units[i].Unitnum)
	}
	pieceNums := fmt.Sprintf("%d",units[0].Piece)
	for i := 1; i < len(units); i++ {
		pieceNums += fmt.Sprintf("%s%d",unitSep, units[i].Piece)
	}
	return unitNums, pieceNums
}

func (s *OPService) CancelOrder(
	ctx context.Context,
	in *grpc_orders.SendOrderReq,
) error {
	return s.repo.CancelOrder(
		ctx, 
		int(in.Userid), 
		int(in.Orderid))
}
