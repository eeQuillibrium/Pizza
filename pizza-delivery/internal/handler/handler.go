package handler

import (
	grpcapp "github.com/eeQuillibrium/pizza-delivery/internal/app/grpc"
	"github.com/eeQuillibrium/pizza-delivery/internal/logger"
	"github.com/eeQuillibrium/pizza-delivery/internal/service"
)

type Handler struct {
	log      *logger.Logger
	services *service.Service
	gRPCApp  *grpcapp.GRPCApp
}

func New(
	log *logger.Logger,
	services *service.Service,
	gRPCApp *grpcapp.GRPCApp,
) *Handler {
	return &Handler{
		log:      log,
		services: services,
		gRPCApp:  gRPCApp,
	}
}
