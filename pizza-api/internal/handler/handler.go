package handler

import (
	grpcapp "github.com/eeQuillibrium/pizza-api/internal/app/grpc"
	"github.com/eeQuillibrium/pizza-api/internal/logger"
	"github.com/eeQuillibrium/pizza-api/internal/service"
	"github.com/gorilla/mux"
)

type Handler struct {
	log     *logger.Logger
	GRPCApp *grpcapp.GRPCApp
}

func New(
	log *logger.Logger,
	authport int,
	kitchenport int,
	kService service.OrderProvider,
) *Handler {
	grpcapp := grpcapp.New(log, authport, kitchenport, kService)
	return &Handler{
		log:     log,
		GRPCApp: grpcapp,
	}
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/orders", h.OrdersHandler)
	r.HandleFunc("/orders/exec", h.OrdersExecHandler)

	r.HandleFunc("/auth/signUp", h.SignUpHandler)
	r.HandleFunc("/auth/signIn", h.SignInHandler)

	return r
}
