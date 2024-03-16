package restapp

import (
	"context"
	"net/http"

	"github.com/eeQuillibrium/pizza-api/internal/app/rest/server"
	"github.com/eeQuillibrium/pizza-api/internal/logger"
)

type RESTApp struct {
	log    *logger.Logger
	port   int
	server *server.GatewayServer
}

func New(
	log *logger.Logger,
	port int,
	handler http.Handler,
) *RESTApp {
	server := server.New(log, port, handler)
	return &RESTApp{
		log:    log,
		port:   port,
		server: server,
	}
}

func (a *RESTApp) Run() {
	a.log.SugaredLogger.Infof("starting rest server on %d", a.port)
	a.server.Run()
}
func (a *RESTApp) Stop(ctx context.Context) {
	a.log.SugaredLogger.Info("stopping rest server")
	if err := a.server.Stop(ctx); err != nil {
		a.log.SugaredLogger.Infof("stop res server error: %w", err)
	}
}
