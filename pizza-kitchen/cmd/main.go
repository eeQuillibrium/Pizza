package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/eeQuillibrium/pizza-kitchen/internal/app"
	"github.com/eeQuillibrium/pizza-kitchen/internal/config"
	"github.com/eeQuillibrium/pizza-kitchen/internal/handler"
	"github.com/eeQuillibrium/pizza-kitchen/internal/logger"
	"github.com/eeQuillibrium/pizza-kitchen/internal/repository"
	"github.com/eeQuillibrium/pizza-kitchen/internal/service"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"

	grpcapp "github.com/eeQuillibrium/pizza-kitchen/internal/app/grpc"
	restapp "github.com/eeQuillibrium/pizza-kitchen/internal/app/rest"
)

func main() {
	log := logger.New()

	if err := godotenv.Load(); err != nil {
		log.Fatalf(".env reading err: %w", err)
	}

	cfg := config.New()

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("localhost:%d", cfg.Repo.Redis.Port),
		Password: cfg.Repo.Redis.Password,
		DB:       cfg.Repo.Redis.DB,
	})

	repo := repository.New(client)
	service := service.New(repo)
	grpcApp := grpcapp.New(
		log, 
		cfg.GRPC.Kitchenapi.Client.Port, 
		cfg.GRPC.Kitchenapi.Server.Port, 
		service,
	)
	handl := handler.New(log, grpcApp, service)
	restApp := restapp.New(log, cfg.REST.Port, handl.InitRoutes())

	app := app.New(grpcApp, restApp)

	go app.Run()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop

	log.SugaredLogger.Infof("try to stop program with %v", sign)

	app.GracefulStop()
}
