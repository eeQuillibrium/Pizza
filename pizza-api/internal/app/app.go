package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	grpcapp "github.com/eeQuillibrium/pizza-api/internal/app/grpc"
	restapp "github.com/eeQuillibrium/pizza-api/internal/app/rest"
	"github.com/eeQuillibrium/pizza-api/internal/config"
	"github.com/eeQuillibrium/pizza-api/internal/handler"
	"github.com/eeQuillibrium/pizza-api/internal/logger"
	"github.com/eeQuillibrium/pizza-api/internal/repository"
	"github.com/eeQuillibrium/pizza-api/internal/service"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type App struct {
	log     *logger.Logger
	RESTApp *restapp.RESTApp
	GRPCApp *grpcapp.GRPCApp
}

func New(
	log *logger.Logger,
) *App {
	cfg := config.New()

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("redis_db:%d", cfg.Repo.Redis.Port),
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

	m, err := migrate.New("migrations/", "postgres://postgres:secret@postgres_db:5433/postgres?sslmode=disable")
	if err != nil {
		log.Fatalf("cannot migrate %v", err)
	}

	if err := m.Up(); err != nil {
		log.Fatalf("migration err %v", err)
	}

	repo := repository.New(log, db, rdb)
	services := service.New(log, repo)
	gRPCApp := grpcapp.New(log, cfg.GRPC.Server.Port, cfg.GRPC.Auth.Port, cfg.GRPC.Kitchen.Port, cfg.GRPC.Delivery.Port, services.OrderProvider)
	handl := handler.New(log, services, gRPCApp)

	restapp := restapp.New(log, cfg.Rest.Port, handl.InitRoutes())
	return &App{
		log:     log,
		RESTApp: restapp,
		GRPCApp: handl.GRPCApp,
	}
}

func (a *App) Run() {
	go a.RESTApp.Run()
	go a.GRPCApp.Run()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stopChan

	a.log.SugaredLogger.Infof("try to shutdown with %v", sign)
	ctx := context.Background()

	a.GracefulStop(ctx)
	/*if err := rdb.ShutdownSave(ctx).Err(); err != nil {
		a.log.SugaredLogger.Infof("error with rdb shutdown : %w", err)
	}*/
}

func (a *App) GracefulStop(ctx context.Context) {
	a.GRPCApp.Stop()
	a.RESTApp.Stop(ctx)
}
