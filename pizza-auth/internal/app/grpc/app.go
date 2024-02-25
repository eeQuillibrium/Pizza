package grpcapp

import (
	"fmt"
	"net"
	"time"

	authgrpc "github.com/eeQuillibrium/pizza-auth/internal/app/server"
	authservice "github.com/eeQuillibrium/pizza-auth/internal/services/auth"
	authstorage "github.com/eeQuillibrium/pizza-auth/internal/storage"
	"github.com/rs/zerolog/log"

	"google.golang.org/grpc"
)

type App struct {
	gRPCServer *grpc.Server
	port       int
	tokenTTL   time.Duration
}

func New(
	port int,
	tokenTTL time.Duration,
) *App {
	gRPCServer := grpc.NewServer()
	authStorage := authstorage.New()
	authService := authservice.New(authStorage, tokenTTL)
	authgrpc.Register(gRPCServer, authService)

	return &App{
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *App) Run() error {
	log.Print("Starting gRPC server...")

	lst, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	log.Print("gRPC server running...")

	if err := a.gRPCServer.Serve(lst); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (a *App) Stop() {
	log.Printf("Stopping gRPC server %v...", a.port)
	a.gRPCServer.GracefulStop() // graceful stopdown
}
