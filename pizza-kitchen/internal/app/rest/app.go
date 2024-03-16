package restapp

import (
	"context"
	"net/http"

	"github.com/eeQuillibrium/pizza-kitchen/internal/app/rest/server"
	"github.com/eeQuillibrium/pizza-kitchen/internal/logger"
)

type RESTApp struct {
	log    *logger.Logger
	Server *server.Server
	port   int
}

func New(
	log *logger.Logger,
	port int,
	router http.Handler,
) *RESTApp {
	serv := server.New(port, router)
	return &RESTApp{
		log:    log,
		Server: serv,
		port:   port,
	}
}

func (a *RESTApp) Run() {
	a.log.SugaredLogger.Infof("try to run rest server on port %d", a.port)
	if err := a.Server.Run(); err != nil {
		a.log.SugaredLogger.Fatalf("server was shutted down")
	}
}

func (a *RESTApp) Stop(ctx context.Context) {
	if err := a.Server.Stop(ctx); err != nil {
		a.log.SugaredLogger.Infof("restserver stopping problem: %w", err)
	}
}
