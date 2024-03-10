package handler

import (
	grpcapp "github.com/eeQuillibrium/pizza-api/internal/app/grpc"
	"github.com/eeQuillibrium/pizza-api/internal/service"
)

type Handler struct {
	GRPCApp *grpcapp.GRPCApp
}

func New(
	authport int,
	kitchenport int,
	kService service.OrderProvider,
) *Handler {
	grpcapp := grpcapp.New(authport, kitchenport, kService)
	return &Handler{grpcapp}
}
