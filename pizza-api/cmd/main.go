package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/eeQuillibrium/pizza-api/internal/app"
	restapp "github.com/eeQuillibrium/pizza-api/internal/app/rest"
	"github.com/eeQuillibrium/pizza-api/internal/config"
	"github.com/eeQuillibrium/pizza-api/internal/handler"
	"github.com/eeQuillibrium/pizza-api/internal/logger"
	"github.com/eeQuillibrium/pizza-api/internal/repository"
	"github.com/eeQuillibrium/pizza-api/internal/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"

	"github.com/joho/godotenv"
)

func main() {
	log := logger.New() //some cfg params

	if err := godotenv.Load(); err != nil {
		log.Fatalf(".env reading err: %v", err)
	}

	cfg := config.New()

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("localhost:%d", cfg.Repo.Redis.Port),
		Password: cfg.Repo.Redis.Password,
		DB:       cfg.Repo.Redis.DB,
	})

	db, err := sqlx.Open("postgres", fmt.Sprintf(
		"user=%s password=%s host=%s dbname=%s port=%d sslmode=%s",
		cfg.Repo.Postgres.Username,
		os.Getenv("DB_PASSWORD"),
		cfg.Repo.Postgres.Host,
		cfg.Repo.Postgres.DBName,
		cfg.Repo.Postgres.Port,
		cfg.Repo.Postgres.SSLMode,
	))
	if err != nil {
		log.SugaredLogger.Fatalf("problem with postgres db.Open() occured: %w", err)
	}
	if err := db.Ping(); err != nil {
		log.SugaredLogger.Fatalf("postgres db ping problem occured: %w", err)
	}

	repo := repository.New(log, db, rdb)
	services := service.New(log, repo)
	handl := handler.New(log, services, cfg.GRPC.Auth.Port, cfg.GRPC.Kitchen.Port, cfg.GRPC.Delivery.Port)

	restapp := restapp.New(log, cfg.Rest.Port, handl.InitRoutes())
	app := app.New(log, restapp, handl.GRPCApp)

	app.Run(cfg.GRPC.Server.Port)

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stopChan

	ctx := context.Background()
	log.SugaredLogger.Infof("try to shutdown with %v", sign)
	app.GracefulStop(ctx)

	if err := rdb.ShutdownSave(ctx).Err(); err != nil {
		log.SugaredLogger.Infof("error with rdb shutdown : %w", err)
	}
}
