package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/eeQuillibrium/pizza-delivery/internal/domain/models"
	grpc_orders "github.com/eeQuillibrium/protos/gen/go/orders"
)

func (h *Handler) ordersGetHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	orders, err := h.services.GetOrders(ctx)
	if err != nil {
		h.log.SugaredLogger.Infof("get orders problem occured: %w", err)
	}

	json, err := json.Marshal(orders)
	if err != nil {
		h.log.SugaredLogger.Fatalf("marshaling problem occured: %w", err)
	}

	w.Write(json)

	h.log.SugaredLogger.Info("successful getorders execution")

}
func (h *Handler) ordersExecHandler(w http.ResponseWriter, r *http.Request) {
	var order models.Order

	jsonBody, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.SugaredLogger.Infof("json reading problem: %w", err)
	}

	json.Unmarshal(jsonBody, &order)

	ctx := context.Background()

	_, err = h.gRPCApp.GatewayClient.SendOrder(
		ctx,
		orderAccessor(&order),
	)

	if err != nil {
		h.log.SugaredLogger.Infof("order sending problem: %w", err)
	}
}
func orderAccessor(order *models.Order) *grpc_orders.SendOrderReq {
	units := []*grpc_orders.SendOrderReq_PieceUnitnum{}
	for i := 0; i < len(order.Units); i++ {
		units = append(units, &grpc_orders.SendOrderReq_PieceUnitnum{
			Piece:   int64(order.Units[i].Piece),
			Unitnum: int64(order.Units[i].Unitnum),
		})
	}

	return &grpc_orders.SendOrderReq{
		Orderid: int64(order.OrderId),
		Userid:  int64(order.UserId),
		Price:   int64(order.Price),
		State:   grpc_orders.SendOrderReq_DELIVERED,
		Units:   units,
	}
}
