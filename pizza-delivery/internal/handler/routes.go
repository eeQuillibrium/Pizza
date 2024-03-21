package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) InitRoutes() http.Handler {
	r := chi.NewRouter()
	r.Get("/", h.homeHandler)

	r.Get("/orders/get", h.ordersGetHandler)
	r.Post("/orders/send/gateway", h.sendGatewayHandler)
	return r
}

func (h *Handler) homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from homepage!"))
}
