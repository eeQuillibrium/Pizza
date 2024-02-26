package app

import (
	"time"

	grpcapp "github.com/eeQuillibrium/pizza-auth/internal/app/grpc"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	port int,
	tokenTTL time.Duration,
	dsn string,
) *App {
	grpcApp := grpcapp.New(port, tokenTTL, dsn)
	return &App{
		GRPCSrv: grpcApp,
	}
}
