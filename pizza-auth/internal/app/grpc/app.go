package grpcapp

import (
	"fmt"
	"net"
	"time"

	authgrpc "github.com/eeQuillibrium/pizza-auth/internal/app/server"
	logger "github.com/eeQuillibrium/pizza-auth/internal/logger"
	authservice "github.com/eeQuillibrium/pizza-auth/internal/services/auth"
	authstorage "github.com/eeQuillibrium/pizza-auth/internal/storage"
	"google.golang.org/grpc"
)

type App struct {
	gRPCServer *grpc.Server
	port       int
	tokenTTL   time.Duration
	log        *logger.Logger
}

func New(
	port int,
	tokenTTL time.Duration,
	dsn string,
	log *logger.Logger,
) *App {
	gRPCServer := grpc.NewServer()
	authStorage := authstorage.New(dsn, log)
	authService := authservice.New(authStorage, tokenTTL, log)
	authgrpc.Register(gRPCServer, authService)

	return &App{
		gRPCServer: gRPCServer,
		port:       port,
		tokenTTL:   tokenTTL,
		log: log,
	}
}

func (a *App) Run() error {
	a.log.SugaredLogger.Info("starting gRPC server...")

	lst, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	a.log.SugaredLogger.Info("gRPC server running...")

	if err := a.gRPCServer.Serve(lst); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (a *App) Stop() {
	a.log.SugaredLogger.Info("stopping gRPC server %v...", a.port)
	a.gRPCServer.GracefulStop() // graceful stopdown
}
