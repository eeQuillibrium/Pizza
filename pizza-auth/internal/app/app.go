package app

import (
	"time"

	grpcapp "github.com/eeQuillibrium/pizza-auth/internal/app/grpc"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	grpcApp := grpcapp.New(grpcPort, tokenTTL)
	return &App{
		GRPCSrv: grpcApp,
	}
}
