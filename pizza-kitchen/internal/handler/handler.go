package handler

import (
	grpcapp "github.com/eeQuillibrium/pizza-kitchen/internal/app/grpc"
	"github.com/eeQuillibrium/pizza-kitchen/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	gRPCApp *grpcapp.GRPCApp
	service *service.Service
}

func New(
	gRPCApp *grpcapp.GRPCApp,
	service *service.Service,
) *Handler {
	return &Handler{
		gRPCApp: gRPCApp,
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/orders/get", h.ordersGetHandler)
	router.GET("/orders/exec", h.ordersExecHandler)

	return router
}
