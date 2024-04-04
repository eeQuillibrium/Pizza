package app

import (
	"time"

	grpcapp "github.com/eeQuillibrium/pizza-auth/internal/app/grpc"
	"github.com/eeQuillibrium/pizza-auth/internal/logger"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	port int,
	tokenTTL time.Duration,
	dsn string,
	log *logger.Logger,
) *App {
	grpcApp := grpcapp.New(port, tokenTTL, dsn, log)
	return &App{
		GRPCSrv: grpcApp,
	}
}
