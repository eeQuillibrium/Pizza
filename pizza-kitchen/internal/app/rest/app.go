package restapp

import (
	"net/http"

	"github.com/eeQuillibrium/pizza-kitchen/internal/app/rest/server"
)

type RESTApp struct {
	Server *server.Server
	port   int
}

func New(
	port int,
	router http.Handler,
) *RESTApp {
	serv := server.New(port, router)
	return &RESTApp{
		Server: serv,
		port:   port,
	}
}

func (a *RESTApp) Run() {
	a.Server.Run()
}
