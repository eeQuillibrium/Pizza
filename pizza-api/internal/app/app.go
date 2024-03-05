package app

import (
	"github.com/eeQuillibrium/pizza-api/internal/app/server"
)

type App struct {
	RESTServ *server.Server
}

func New(
	RESTServ *server.Server,
) *App {
	return &App{RESTServ}
}
