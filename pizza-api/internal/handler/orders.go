package handler

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/eeQuillibrium/pizza-api/internal/domain/models"
	nikita_kitchen1 "github.com/eeQuillibrium/protos/gen/go/kitchen"
)

func (h *Handler) OrderHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Order request, method:", r.Method+"...")
}

func (h *Handler) OrderExecHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("orderExec requested with method:", r.Method, r.Body)

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("order json reading problem: %v", err)
	}

	log.Printf("req body:%s", string(b))

	if len(b) == 0 {
		log.Fatal("empty req body")
	}

	var order models.Order
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

	h.GRPCApp.Kitchen.SendOrder(
		ctx,
		&nikita_kitchen1.SendOrderReq{Units: units, Userid: int64(order.UserId), Price: int64(order.Price)},
	)

	log.Print("successful sendorder execution")
}
