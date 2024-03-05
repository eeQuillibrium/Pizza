package handler

import grpcapp "github.com/eeQuillibrium/pizza-kitchen/internal/app/grpc"

type Handler struct {
	GRPCApp *grpcapp.GRPCApp
}

func New(
	GRPCApp *grpcapp.GRPCApp,
) *Handler {
	return &Handler{
		GRPCApp,
	}
}

func InitRoutes() {

}
