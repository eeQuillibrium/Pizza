package app

import (
	"context"

	grpcapp "github.com/eeQuillibrium/pizza-delivery/internal/app/grpc"
	restapp "github.com/eeQuillibrium/pizza-delivery/internal/app/rest"
)

type App struct {
	RESTApp *restapp.RESTApp
	GRPCApp *grpcapp.GRPCApp
}

func New(
	RESTApp *restapp.RESTApp,
	GRPCApp *grpcapp.GRPCApp,
) *App {
	return &App{
		RESTApp: RESTApp,
		GRPCApp: GRPCApp,
	}
}
func (a *App) Run(gatewayServerPort int) {
	go a.RESTApp.Run()
	a.GRPCApp.Run(gatewayServerPort)
}
func (a *App) GracefulStop(ctx context.Context) {
	a.RESTApp.Stop(ctx)
	a.GRPCApp.Stop()
}
