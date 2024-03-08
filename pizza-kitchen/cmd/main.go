package main

import (
	"fmt"
	"log"

	"github.com/eeQuillibrium/pizza-kitchen/internal/app"
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
	handl := handler.New(service)
	router := handl.InitRoutes()

	app := app.New(cfg.GRPC.Kitchenapi.Port, cfg.REST.Port, router, service)

	app.Run()

	//Graceful shutdown
}
