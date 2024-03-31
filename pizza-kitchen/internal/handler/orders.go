package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	grpc_orders "github.com/eeQuillibrium/protos/gen/go/orders"
	"github.com/gin-gonic/gin"

	"github.com/eeQuillibrium/pizza-kitchen/internal/domain/models"
)

func (h *Handler) ordersGetHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	orders, err := h.service.Kitchen.GetOrders(ctx)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		h.log.SugaredLogger.Infof("getorders problem: %w", err)
	}

	json, err := json.Marshal(orders)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		h.log.SugaredLogger.Infof("json marshaling problem occured: %w", err)
	}

	c.Writer.Write(json)
}

func (h *Handler) sendDeliveryHandler(c *gin.Context) {
	b, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		h.log.SugaredLogger.Fatalf("order json reading problem: %w", err)
	}
	if len(b) == 0 {
		c.Status(http.StatusInternalServerError)
		h.log.SugaredLogger.Fatal("empty req body")
	}

	order := models.Order{}

	json.Unmarshal(b, &order)

	ctx := context.Background()

	if err := h.service.DeleteOrder(ctx, order.OrderId); err != nil {
		c.Status(http.StatusInternalServerError)
		h.log.SugaredLogger.Infof("deleting problem occured: %w", err)
	}

	gReq := orderAccessor(&order)

	if _, err = h.gRPCApp.GatewayOS.SendOrder(ctx, gReq); err != nil {
		c.Status(http.StatusBadGateway)
		h.log.SugaredLogger.Fatalf("error with sendorder to gateway: %v", err)
	}

	if _, err = h.gRPCApp.DeliveryOS.SendOrder(ctx, gReq); err != nil {
		c.Status(http.StatusBadGateway)
		h.log.SugaredLogger.Fatalf("error with sendorder to delivery: %v", err)
	}

	c.Status(http.StatusOK)
}

func (h *Handler) ordersCancel(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	order := models.Order{}

	json.Unmarshal(data, &order)

	order.State = "CANCELLED"
	grpcOrder := orderAccessor(&order)

	if _, err = h.gRPCApp.GatewayOS.SendOrder(ctx, grpcOrder); err != nil {
		c.Status(http.StatusBadGateway)
		h.log.SugaredLogger.Fatalf("error with sendorder to gateway: %v", err)
	}

	if _, err = h.gRPCApp.DeliveryOS.SendOrder(ctx, grpcOrder); err != nil {
		c.Status(http.StatusBadGateway)
		h.log.SugaredLogger.Fatalf("error with sendorder to delivery: %v", err)
	}

	if err := h.service.Kitchen.CancelOrder(ctx, &order); err != nil {
		c.Status(http.StatusInternalServerError)
		h.log.SugaredLogger.Fatalf("error with cancel order: %v", err)
	}

	c.Status(http.StatusOK)
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
		State:   grpc_orders.SendOrderReq_DELIVER,
	}
}
