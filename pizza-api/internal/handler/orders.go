package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/eeQuillibrium/pizza-api/internal/domain/models"
	grpc_orders "github.com/eeQuillibrium/protos/gen/go/orders"
)

func (h *Handler) ordersGetHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	orders, err := h.service.APIProvider.GetOrders(ctx)
	if err != nil {
		h.log.SugaredLogger.Fatalf("getting order problem %w", err)
	}

	json, err := json.Marshal(orders)
	if err != nil {
		h.log.SugaredLogger.Fatalf("marshaling problem %w", err)
	}

	w.Write(json)

	h.log.SugaredLogger.Info("successful getorders execution")
}

func (h *Handler) ordersExecHandler(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.SugaredLogger.Fatalf("order json reading problem: %w", err)
	}
	if len(b) == 0 {
		h.log.SugaredLogger.Fatal("empty req body")
	}

	var order models.Order
	json.Unmarshal(b, &order)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = h.GRPCApp.KitchenOrderSender.SendOrder(
		ctx,
		&grpc_orders.SendOrderReq{
			Units:  orderUnitsAccessor(&order),
			Userid: int64(order.UserId),
			Price:  int64(order.Price),
		},
	)
	if err != nil {
		h.log.SugaredLogger.Fatalf("error with sendorder: %w", err)
	}

	h.log.SugaredLogger.Info("successful order storing in kitchen")
}

func orderUnitsAccessor(order *models.Order) []*grpc_orders.SendOrderReq_PieceUnitnum {
	units := []*grpc_orders.SendOrderReq_PieceUnitnum{}
	for i := 0; i < len(order.Units); i++ {
		units = append(units, &grpc_orders.SendOrderReq_PieceUnitnum{
			Piece:   int64(order.Units[i].Piece),
			Unitnum: int64(order.Units[i].Unitnum),
		})
	}
	return units
}
