package handler

import (
	grpcapp "github.com/eeQuillibrium/pizza-kitchen/internal/app/grpc"
	"github.com/eeQuillibrium/pizza-kitchen/internal/logger"
	"github.com/eeQuillibrium/pizza-kitchen/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	log     *logger.Logger
	gRPCApp *grpcapp.GRPCApp
	service *service.Service
}

func New(
	log *logger.Logger,
	gRPCApp *grpcapp.GRPCApp,
	service *service.Service,
) *Handler {
	return &Handler{
		log:     log,
		gRPCApp: gRPCApp,
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/orders")
	router.GET("/orders/get", h.ordersGetHandler)
	router.DELETE("/orders/cancel", h.ordersCancel)
	router.POST("/orders/send/delivery", h.sendDeliveryHandler)

	return router
}
