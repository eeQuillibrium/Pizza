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
	gRPCApp *grpcapp.GRPCApp,
) *Handler {

	return &Handler{
		log:     log,
		service: service,
		GRPCApp: gRPCApp,
	}
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/auth", h.authHandler)
	auth.HandleFunc("/auth/signUp", h.signUpHandler)
	auth.HandleFunc("/auth/signIn", h.signInHandler)

	home := r.PathPrefix("/home").Subrouter()
	home.HandleFunc("/", h.homeHandler)

	orders := r.PathPrefix("/orders").Subrouter()
	orders.HandleFunc("/", h.ordersHandler)
	orders.HandleFunc("/create", h.createOrderHandler)
	orders.HandleFunc("/cancel", h.ordersCancelHandler)
	orders.HandleFunc("/current", h.ordersCurrentHandler)
	orders.HandleFunc("/history", h.ordersHistoryHandler)
	orders.Use(h.userIdentify)

	review := r.PathPrefix("/review").Subrouter()
	review.HandleFunc("/", h.reviewHandler)
	review.HandleFunc("/send", h.reviewSendHandler)
	review.Use(h.userIdentify)

	return r
}
