package handler

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"time"

	nikita_kitchen1 "github.com/eeQuillibrium/protos/gen/go/kitchen"

	"github.com/eeQuillibrium/pizza-kitchen/internal/domain/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/", h.homepageHandler)
	return router
}

func (h *Handler) homepageHandler(c *gin.Context) {
	log.Print("homepage kitchen was GETTED")

	var order models.Order
	b, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatalf("order json decode problem: %v", err)
	}
	log.Print(string(b))
	json.Unmarshal(b, &order)

	var units []*nikita_kitchen1.SendOrderReq_PieceUnitnum = make([]*nikita_kitchen1.SendOrderReq_PieceUnitnum, len(order.Units))
	for i := 0; i < len(order.Units); i++ {
		units = append(units, &nikita_kitchen1.SendOrderReq_PieceUnitnum{
			Piece:   int64(order.Units[i].Piece),
			Unitnum: int64(order.Units[i].Unitnum),
		})
	}

	log.Printf("order parameters: UserId: %d Price: %d Units: %v", order.UserId, order.Price, order.Units)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = h.gRPCApp.OrderSender.SendOrder(
		ctx,
		&nikita_kitchen1.SendOrderReq{Units: units, Userid: int64(order.UserId), Price: int64(order.Price)},
	)

	if err != nil {
		log.Fatalf("err with sendmessage: %v", err)
	}

	log.Print("successful sendmessage execution")
}
