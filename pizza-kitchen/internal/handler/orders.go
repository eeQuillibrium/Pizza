package handler

import (
	"context"
	"encoding/json"
	"io"
	"time"

	grpc_orders "github.com/eeQuillibrium/protos/gen/go/orders"
	"github.com/gin-gonic/gin"

	"github.com/eeQuillibrium/pizza-kitchen/internal/domain/models"
)

func (h *Handler) ordersGetHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	orders, err := h.service.Kitchen.GetOrders(ctx)
	if err != nil {
		h.log.SugaredLogger.Infof("getorders problem: %w", err)
	}

	json, err := json.Marshal(orders)
	if err != nil {
		h.log.SugaredLogger.Infof("json marshaling problem occured: %w", err)
	}

	c.Writer.Write(json)
	h.log.SugaredLogger.Info("successful getorders execution")
}
func (h *Handler) ordersExecHandler(c *gin.Context) {
	b, err := io.ReadAll(c.Request.Body)
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

	_, err = h.gRPCApp.APIOrderSender.SendOrder(
		ctx,
		&grpc_orders.SendOrderReq{
			Units:  orderUnitsAccessor(&order),
			Userid: int64(order.UserId),
			Price:  int64(order.Price),
		},
	)
	if err != nil {
		h.log.SugaredLogger.Fatalf("error with sendorder: %v", err)
	}

	h.log.SugaredLogger.Info("successful order storing in gateway!")
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
