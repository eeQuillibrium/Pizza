package app

import (
	"github.com/eeQuillibrium/pizza-api/internal/app/server"
	"github.com/gorilla/mux"
)

type App struct {
	RESTServ *server.Server
}

func New(
	restport string,
	r *mux.Router,
) *App {
	RESTServ := server.New(restport, r)
	return &App{RESTServ}
}
