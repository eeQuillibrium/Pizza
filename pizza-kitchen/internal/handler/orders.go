package handler

import (
	"context"
	"encoding/json"
	"io"
	"log"
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
		log.Printf("getorders problem: %v", err)
	}

	jsonString, err := json.Marshal(orders)
	if err != nil {
		log.Printf("json marshaling problem occured: %v", err)
	}

	c.Writer.Write(jsonString)
}
func (h *Handler) ordersExecHandler(c *gin.Context) {
	b, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatalf("order json reading problem: %v", err)
	}
	if len(b) == 0 {
		log.Fatal("empty req body")
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
		log.Fatalf("error with sendorder: %v", err)
	}

	log.Print("successful sendorder execution")
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
