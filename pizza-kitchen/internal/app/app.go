package app

import (
	"net/http"

	grpcapp "github.com/eeQuillibrium/pizza-kitchen/internal/app/grpc"
	restapp "github.com/eeQuillibrium/pizza-kitchen/internal/app/rest"
	"github.com/eeQuillibrium/pizza-kitchen/internal/service"
)

type App struct {
	GRPCApp  *grpcapp.GRPCApp
	RESTServ *restapp.RESTApp
}

func New(
	grpcPortApi int,
	//grpcPortDel int,
	restport int,
	router http.Handler,
	service *service.Service,
) *App {

	return &App{
		grpcapp.New(grpcPortApi, service),
		restapp.New(restport, router),
	}
}

func (a *App) Run() {

	a.GRPCApp.Run()
	a.RESTServ.Run()
}
