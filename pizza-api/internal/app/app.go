package app

import (
	"context"

	grpcapp "github.com/eeQuillibrium/pizza-api/internal/app/grpc"
	"github.com/eeQuillibrium/pizza-api/internal/app/server"
	"github.com/eeQuillibrium/pizza-api/internal/logger"
)

type App struct {
	log      *logger.Logger
	RESTServ *server.Server
	GRPCApp  *grpcapp.GRPCApp
}

func New(
	log *logger.Logger,
	RESTServ *server.Server,
	GRPCApp *grpcapp.GRPCApp,
) *App {
	return &App{
		log:      log,
		RESTServ: RESTServ,
		GRPCApp:  GRPCApp,
	}
}

func (a *App) Run(orderPort int) {
	go a.RESTServ.Run()
	go a.GRPCApp.Run(orderPort)
}
func (a *App) GracefulStop(ctx context.Context) {
	a.GRPCApp.Stop()
	if err := a.RESTServ.Stop(ctx); err != nil {
		a.log.SugaredLogger.Infof("stop res serv error: %w", err)
	}
}
