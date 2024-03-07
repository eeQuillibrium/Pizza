package main

import (
	"log"

	"github.com/eeQuillibrium/pizza-kitchen/internal/app"
	"github.com/eeQuillibrium/pizza-kitchen/internal/config"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf(".env reading err: %v", err)
	}

	log.Print("kitchen microservice start...")

	cfg := config.New()

	app := app.New(cfg.GRPC.Kitchenapi.Port, cfg.REST.Port)

	app.Run()

	//Graceful shutdown
}
