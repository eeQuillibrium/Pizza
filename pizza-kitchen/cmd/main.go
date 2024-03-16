package main

import (
	"context"
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

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("localhost:%d", cfg.Repo.Redis.Port),
		Password: cfg.Repo.Redis.Password,
		DB:       cfg.Repo.Redis.DB,
	})

	repo := repository.New(rdb)
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

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stopChan

	ctx := context.Background()
	log.SugaredLogger.Infof("try to stop program with %v", sign)
	app.GracefulStop(ctx)

	if err := rdb.ShutdownSave(ctx).Err(); err != nil {
		log.SugaredLogger.Infof("error with rdb shutdown : %w", err)
	}
}
