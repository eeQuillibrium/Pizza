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
	
}
func (h *Handler) ordersGetHandler(w http.ResponseWriter, r *http.Request) {
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	userId := struct{UserId int`json:"userid"`}{}
	json.Unmarshal(data, &userId)

	orders, err := h.service.GetCurrentOrders(ctx, userId.UserId)
	if err != nil {
		h.log.SugaredLogger.Fatalf("getting order problem %w", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	json, err := json.Marshal(orders)
	if err != nil {
		h.log.SugaredLogger.Fatalf("marshaling problem %w", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write(json)

	h.log.SugaredLogger.Info("successful getorders execution")
}

func (h *Handler) sendKitchenHandler(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.SugaredLogger.Fatalf("order json reading problem: %w", err)
	}
	if len(b) == 0 {
		h.log.SugaredLogger.Fatal("empty req body")
	}

	var order models.Order
	json.Unmarshal(b, &order)

	ctx := context.Background()

	h.service.APIProvider.CreateOrder(ctx, &order)

	if _, err = h.GRPCApp.KitchenOrderSender.SendOrder(
		ctx,
		orderAccessor(&order),
	); err != nil {
		h.log.SugaredLogger.Fatalf("error with sendorder: %w", err)
	}

	h.log.SugaredLogger.Info("successful order storing in kitchen")
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

func (h *Handler) homeHandler(w http.ResponseWriter, r *http.Request) {

}
