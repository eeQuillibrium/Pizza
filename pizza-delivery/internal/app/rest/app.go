package restapp

import (
	"context"
	"net/http"

	"github.com/eeQuillibrium/pizza-delivery/internal/app/rest/server"
	"github.com/eeQuillibrium/pizza-delivery/internal/logger"
)

type RESTApp struct {
	log         *logger.Logger
	server      *server.GatewayServer
	gatewayPort int
}

func New(
	log *logger.Logger,
	gatewayPort int,
	router http.Handler,
) *RESTApp {
	server := server.New(gatewayPort, router)
	return &RESTApp{
		log:         log,
		server:      server,
		gatewayPort: gatewayPort,
	}
}

func (a *RESTApp) Run() {
	a.log.SugaredLogger.Infof("run rest server on %d", a.gatewayPort)
	if err := a.server.Run(); err != nil {
		a.log.SugaredLogger.Infof("listen problem: %w", err)
	}
}

func (a *RESTApp) Stop(ctx context.Context) {
	a.log.SugaredLogger.Info("stopping rest server")
	if err := a.server.Shutdown(ctx); err != nil {
		a.log.SugaredLogger.Infof("stopping rest error %w", err)
	}
}
