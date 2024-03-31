package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/eeQuillibrium/pizza-delivery/internal/app"
	grpcapp "github.com/eeQuillibrium/pizza-delivery/internal/app/grpc"
	restapp "github.com/eeQuillibrium/pizza-delivery/internal/app/rest"
	"github.com/eeQuillibrium/pizza-delivery/internal/config"
	"github.com/eeQuillibrium/pizza-delivery/internal/handler"
	"github.com/eeQuillibrium/pizza-delivery/internal/logger"
	"github.com/eeQuillibrium/pizza-delivery/internal/repository"
	"github.com/eeQuillibrium/pizza-delivery/internal/service"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func main() {
	log := logger.New()
	if err := godotenv.Load(); err != nil {
		log.SugaredLogger.Fatalf("load .env error: %w", err)
	}

	cfg := config.New(log)

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("localhost:%d", cfg.Repo.Redis.Port),
		Password: cfg.Repo.Redis.Password,
		DB:       cfg.Repo.Redis.DB,
	})

	repo := repository.New(rdb)
	services := service.New(repo)

	gRPCApp := grpcapp.New(log, services.OrderProvider, cfg.GRPC.GatewayClient.Port, cfg.GRPC.KitchenClient.Port)
	handl := handler.New(log, services, gRPCApp)
	restApp := restapp.New(log, cfg.REST.Port, handl.InitRoutes())

	appl := app.New(restApp, gRPCApp)

	go appl.Run(cfg.GRPC.Server.Port)

	signChan := make(chan os.Signal, 1)
	signal.Notify(signChan, syscall.SIGTERM, syscall.SIGINT)

	sign := <-signChan

	log.SugaredLogger.Infof("try to stop app with %v", sign)

	ctx := context.Background()

	appl.GracefulStop(ctx)

	if err := rdb.ShutdownSave(ctx).Err(); err != nil {
		log.SugaredLogger.Infof("error with rdb shutdown : %w", err)
	}
}
