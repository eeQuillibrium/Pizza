package main

import (
	"fmt"

	"github.com/eeQuillibrium/pizza-api/internal/app"
	"github.com/eeQuillibrium/pizza-api/internal/app/server"
	"github.com/eeQuillibrium/pizza-api/internal/config"
	"github.com/eeQuillibrium/pizza-api/internal/handler"
	"github.com/eeQuillibrium/pizza-api/internal/logger"
	"github.com/eeQuillibrium/pizza-api/internal/repository"
	"github.com/eeQuillibrium/pizza-api/internal/service"
	"github.com/redis/go-redis/v9"

	"github.com/joho/godotenv"
)

func main() {
	log := logger.New() //some cfg params

	if err := godotenv.Load(); err != nil {
		log.Fatalf(".env reading err: %v", err)
	}

	cfg := config.New()

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("localhost:%d", cfg.Repo.Redis.Port),
		Password: cfg.Repo.Redis.Password,
		DB:       cfg.Repo.Redis.DB,
	})

	repo := repository.New(log, client)
	services := service.New(log, repo)
	handl := handler.New(log, services, cfg.GRPC.Auth.Port, cfg.GRPC.Kitchen.Port, services)

	RESTServ := server.New(log, cfg.Server.Port, handl.InitRoutes())

	app := app.New(RESTServ, handl.GRPCApp)

	app.Run(cfg.GRPC.KitchenOrder.Port)

	//graceful shutdown
}
