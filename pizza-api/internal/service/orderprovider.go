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
	unitNums := ""
	for i := 0; i < len(units)-1; i++ {
		unitNums += fmt.Sprintf("%d,", units[i].Unitnum)
	}
	pieceNums := ""
	for i := 0; i < len(units)-1; i++ {
		pieceNums += fmt.Sprintf("%d,", units[i].Piece)
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
