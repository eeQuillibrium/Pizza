package app

import (
	grpcapp "github.com/eeQuillibrium/pizza-api/internal/app/grpc"
	"github.com/eeQuillibrium/pizza-api/internal/app/server"
)

type App struct {
	RESTServ *server.Server
	GRPCApp  *grpcapp.GRPCApp
}

func New(
	RESTServ *server.Server,
	GRPCApp *grpcapp.GRPCApp,
) *App {
	return &App{
		RESTServ: RESTServ,
		GRPCApp:  GRPCApp,
	}
}

func (a *App) Run(orderPort int) {
	go a.RESTServ.Run()
	a.GRPCApp.Run(orderPort)
}
