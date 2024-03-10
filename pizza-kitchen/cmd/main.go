package main

import (
	"fmt"
	"log"

	"github.com/eeQuillibrium/pizza-kitchen/internal/app"
	grpcapp "github.com/eeQuillibrium/pizza-kitchen/internal/app/grpc"
	restapp "github.com/eeQuillibrium/pizza-kitchen/internal/app/rest"
	"github.com/eeQuillibrium/pizza-kitchen/internal/config"
	"github.com/eeQuillibrium/pizza-kitchen/internal/handler"
	"github.com/eeQuillibrium/pizza-kitchen/internal/repository"
	"github.com/eeQuillibrium/pizza-kitchen/internal/service"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf(".env reading err: %v", err)
	}

	log.Print("try to start...")

	cfg := config.New()

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("localhost:%d", cfg.Repo.Redis.Port),
		Password: cfg.Repo.Redis.Password,
		DB:       cfg.Repo.Redis.DB,
	})

	repo := repository.New(client)
	service := service.New(repo)

	grpcApp := grpcapp.New(cfg.GRPC.Kitchenapi.Client.Port, cfg.GRPC.Kitchenapi.Server.Port, service)
	handl := handler.New(grpcApp, service)
	restApp := restapp.New(cfg.REST.Port, handl.InitRoutes())

	app := app.New(grpcApp, restApp)

	app.Run()

	//Graceful shutdown
}
