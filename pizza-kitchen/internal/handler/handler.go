package handler

import (
	grpcapp "github.com/eeQuillibrium/pizza-kitchen/internal/app/grpc"
	"github.com/eeQuillibrium/pizza-kitchen/internal/service"
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
