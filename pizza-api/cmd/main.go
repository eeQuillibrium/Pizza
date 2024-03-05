package main

import (
	"log"

	"github.com/eeQuillibrium/pizza-api/internal/app"
	"github.com/eeQuillibrium/pizza-api/internal/app/server"
	"github.com/eeQuillibrium/pizza-api/internal/config"
	"github.com/eeQuillibrium/pizza-api/internal/handler"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf(".env reading err: %v", err)
	}

	cfg := config.New()

	handl := handler.New(cfg.GRPC.Auth.Port, cfg.GRPC.Kitchen.Port)

	RESTServ := server.New(cfg.Server.Port, handl.InitRoutes())
	app := app.New(RESTServ)

	app.RESTServ.Run()

	//graceful shutdown
}
