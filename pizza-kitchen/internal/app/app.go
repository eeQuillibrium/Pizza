package app

import (
	grpcapp "github.com/eeQuillibrium/pizza-kitchen/internal/app/grpc"
	restapp "github.com/eeQuillibrium/pizza-kitchen/internal/app/rest"
	handler "github.com/eeQuillibrium/pizza-kitchen/internal/handlers"
)

type App struct {
	GRPCApp  *grpcapp.GRPCApp
	RESTServ *restapp.RESTApp
}

func New(
	grpcPortApi int,
	//grpcPortDel int,
	restport int,
) *App {
	router := handler.InitRoutes()
	
	return &App{
		grpcapp.New(grpcPortApi),
		restapp.New(restport, router),
	}
}

func (a *App) Run() {
	a.GRPCApp.Run()
	a.RESTServ.Run()
}
