package app

import (
	"github.com/eeQuillibrium/pizza-kitchen/internal/app/rest"
	"github.com/eeQuillibrium/pizza-kitchen/internal/app/grpc"
)

type App struct {
	GRPCApp  *grpcapp.GRPCApp
	RESTServ *restapp.RESTApp
}

func New() *App {
	return &App{
		grpcapp.New(),
		restapp.New(),
	}
}