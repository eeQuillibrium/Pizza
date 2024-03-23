package handler

import (
	grpcapp "github.com/eeQuillibrium/pizza-api/internal/app/grpc"
	"github.com/eeQuillibrium/pizza-api/internal/logger"
	"github.com/eeQuillibrium/pizza-api/internal/service"
	"github.com/gorilla/mux"
)

type Handler struct {
	log     *logger.Logger
	service *service.Service
	GRPCApp *grpcapp.GRPCApp
}

func New(
	log *logger.Logger,
	service *service.Service,
	authport int,
	kitchenport int,
	kService service.OrderProvider,
) *Handler {
	grpcapp := grpcapp.New(log, authport, kitchenport, kService)
	return &Handler{
		log:     log,
		service: service,
		GRPCApp: grpcapp,
	}
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/home", h.homeHandler)

	r.HandleFunc("/orders", h.ordersHandler)
	r.HandleFunc("/orders/send/kitchen", h.sendKitchenHandler)
	r.HandleFunc("/orders/current", h.ordersCurrentHandler)
	r.HandleFunc("/orders/history", h.ordersHistoryHandler)

	r.HandleFunc("/auth", h.authHandler)
	r.HandleFunc("/auth/signUp", h.signUpHandler)
	r.HandleFunc("/auth/signIn", h.signInHandler)

	r.HandleFunc("/review", h.reviewHandler)
	r.HandleFunc("/review/send", h.reviewSendHandler)

	return r
}
