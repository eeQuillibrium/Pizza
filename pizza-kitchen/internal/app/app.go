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
	go a.RESTServ.Run()
}

func (a *App) GracefulStop() {
	a.GRPCApp.Stop()
	a.RESTServ.Stop()
}