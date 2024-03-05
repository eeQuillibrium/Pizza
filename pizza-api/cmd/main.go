package main

import (
	"log"

	"github.com/eeQuillibrium/pizza-api/internal/app"
	"github.com/eeQuillibrium/pizza-api/internal/config"
	"github.com/eeQuillibrium/pizza-api/internal/handler"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf(".env reading err: %v", err)
	}

	cfg := config.New()

	handl := handler.New(cfg.GRPC.Port)

	app := app.New(cfg.Server.Port, handl.InitRoutes())

	app.RESTServ.Run("grpc.port")

	//graceful shutdown
}
