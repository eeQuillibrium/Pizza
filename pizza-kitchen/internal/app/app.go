package app

import (
	grpcapp "github.com/eeQuillibrium/pizza-kitchen/internal/app/grpc"
	restapp "github.com/eeQuillibrium/pizza-kitchen/internal/app/rest"
)

type App struct {
	GRPCApp  *grpcapp.GRPCApp
	RESTServ *restapp.RESTApp
}

func New(
	GRPCApp *grpcapp.GRPCApp,
	RESTServ *restapp.RESTApp,
) *App {

	return &App{
		GRPCApp:  GRPCApp,
		RESTServ: RESTServ,
	}
}

func (a *App) Run() {
	go a.GRPCApp.Run()
	a.RESTServ.Run()
}
