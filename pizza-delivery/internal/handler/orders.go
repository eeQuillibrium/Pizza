package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/eeQuillibrium/pizza-delivery/internal/domain/models"
	grpc_orders "github.com/eeQuillibrium/protos/gen/go/orders"
)

func (h *Handler) ordersHandler(w http.ResponseWriter, r *http.Request) {
}
func (h *Handler) ordersCancelHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.log.SugaredLogger.Fatalf("error in body readeing: %w", err)
	}

	order := models.Order{}
	json.Unmarshal(data, &order)

	in := orderAccessor(&order)

	if _, err := h.gRPCApp.GatewayClient.SendOrder(ctx, in); err != nil {
		w.WriteHeader(http.StatusBadGateway)
		h.log.SugaredLogger.Fatalf("gateway sending problem: %w", err)
	}

	if _, err := h.gRPCApp.KitchenClient.SendOrder(ctx, in); err != nil {
		w.WriteHeader(http.StatusBadGateway)
		h.log.SugaredLogger.Fatalf("kitchen sending problem: %w", err)
	}

	if err := h.services.APIProvider.DeleteOrder(ctx, order.OrderId); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.log.SugaredLogger.Fatalf("cancel order problem: %w", err)
	}

	w.WriteHeader(http.StatusOK)
}
func (h *Handler) ordersGetHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	orders, err := h.services.GetCurrentOrders(ctx)
	if err != nil {
		h.log.SugaredLogger.Infof("get orders problem occured: %w", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	json, err := json.Marshal(orders)
	if err != nil {
		h.log.SugaredLogger.Fatalf("marshaling problem occured: %w", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write(json)
}
func (h *Handler) sendGatewayHandler(w http.ResponseWriter, r *http.Request) {
	var order models.Order

	jsonBody, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.SugaredLogger.Infof("json reading problem: %w", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.Unmarshal(jsonBody, &order)

	ctx := context.Background()

	if _, err = h.gRPCApp.GatewayClient.SendOrder(
		ctx,
		orderAccessor(&order),
	); err != nil {
		h.log.SugaredLogger.Infof("order sending problem: %w", err)
		w.WriteHeader(http.StatusBadGateway)
	}

	if err := h.services.APIProvider.DeleteOrder(ctx, order.OrderId); err != nil {
		h.log.SugaredLogger.Infof("order deleting problem: %w", err)
		w.WriteHeader(http.StatusInternalServerError)
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
