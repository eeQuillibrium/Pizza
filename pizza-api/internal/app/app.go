package app

import (
	"context"

	grpcapp "github.com/eeQuillibrium/pizza-api/internal/app/grpc"
	restapp "github.com/eeQuillibrium/pizza-api/internal/app/rest"
	"github.com/eeQuillibrium/pizza-api/internal/logger"
)

type App struct {
	log     *logger.Logger
	RESTApp *restapp.RESTApp
	GRPCApp *grpcapp.GRPCApp
}

func New(
	log *logger.Logger,
	RESTApp *restapp.RESTApp,
	GRPCApp *grpcapp.GRPCApp,
) *App {
	return &App{
		log:     log,
		RESTApp: RESTApp,
		GRPCApp: GRPCApp,
	}
}

func (a *App) Run(orderPort int) {
	go a.RESTApp.Run()
	go a.GRPCApp.Run(orderPort)
}
func (a *App) GracefulStop(ctx context.Context) {
	a.GRPCApp.Stop()
	a.RESTApp.Stop(ctx)
}
