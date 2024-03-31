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

func (h *Handler) ordersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from /orders/ route"))
}

func (h *Handler) createOrderHandler(w http.ResponseWriter, r *http.Request) {

	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.log.SugaredLogger.Fatalf("order json reading problem: %w", err)
	}
	if len(b) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		h.log.SugaredLogger.Fatal("empty req body")
	}

	var order models.Order
	if err := json.Unmarshal(b, &order); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.log.SugaredLogger.Fatalf("order json reading problem: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	order.State = grpc_orders.SendOrderReq_ORDERED.String()
	if err := h.service.APIProvider.CreateOrder(ctx, &order); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.log.SugaredLogger.Fatalf("error with sendorder: %w", err)
	}

	order.State = grpc_orders.SendOrderReq_COOK.String()
	if _, err := h.GRPCApp.KitchenOS.SendOrder(ctx, orderAccessor(&order)); err != nil {
		w.WriteHeader(http.StatusBadGateway)
		h.log.SugaredLogger.Fatalf("error with sendorder: %w", err)
	}
}
func (h *Handler) ordersCancelHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.log.SugaredLogger.Fatalf("body readeing problem occured: %w", err)
	}

	order := models.Order{}
	json.Unmarshal(data, &order)

	grpcOrder := orderAccessor(&order)
	grpcOrder.State = grpc_orders.SendOrderReq_CANCELLED
	if _, err := h.GRPCApp.KitchenOS.SendOrder(ctx, grpcOrder); err != nil {
		w.WriteHeader(http.StatusBadGateway)
		h.log.SugaredLogger.Fatalf("error kitchen order deleting: %v", err)
	}

	if _, err := h.GRPCApp.DeliveryOS.SendOrder(ctx, grpcOrder); err != nil {
		w.WriteHeader(http.StatusBadGateway)
		h.log.SugaredLogger.Fatalf("error delivery order deleting: %v", err)
	}

	if err := h.service.APIProvider.CancelOrder(ctx, &order); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.log.SugaredLogger.Fatalf("error int order deleting: %w", err)
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) ordersCurrentHandler(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	userId := struct {
		UserId int `json:"userid"`
	}{}
	json.Unmarshal(data, &userId)

	orders, err := h.service.GetCurrentOrders(ctx, userId.UserId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.log.SugaredLogger.Fatalf("getting order problem %w", err)
	}

	json, err := json.Marshal(orders)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.log.SugaredLogger.Fatalf("marshaling problem %w", err)
	}

	w.Write(json)

	h.log.SugaredLogger.Info("successful getorders execution")
}

func (h *Handler) ordersHistoryHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	userId := struct {
		UserId int `json:"userid"`
	}{}
	json.Unmarshal(data, &userId)
	h.log.SugaredLogger.Infof("userid: %d", userId.UserId)
	orders, err := h.service.APIProvider.GetOrdersHistory(ctx, userId.UserId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.log.SugaredLogger.Fatalf("error with get history: %w", err)
	}

	json, err := json.Marshal(orders)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.log.SugaredLogger.Fatalf("marshaling problem %w", err)
	}

	w.Write(json)
}

func (h *Handler) homeHandler(w http.ResponseWriter, r *http.Request) {

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
		Units:   units,
		State:   grpc_orders.SendOrderReq_COOK, //*
	}
}
