package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/eeQuillibrium/pizza-auth/internal/app"
	"github.com/eeQuillibrium/pizza-auth/internal/config"
	"github.com/eeQuillibrium/pizza-auth/internal/logger"
	"github.com/joho/godotenv"
)

func main() {

	log := logger.New()

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Env var didn't loaded: %v", err)
	}

	cfg := config.New(log)

	appl := app.New(
		cfg.GRPC.Port,
		cfg.TokenTTL,
		fmt.Sprintf(
			"user=%s password=%s host=%s dbname=%s port=%d sslmode=%s",
			cfg.Storage.Username,
			os.Getenv("DB_PASSWORD"),
			cfg.Storage.Host,
			cfg.Storage.DBName,
			cfg.Storage.Port,
			cfg.Storage.SSLMode,
		),
		log,
	)

	go appl.GRPCSrv.Run()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop

	log.Infof("try to stop the program with ", sign)
	appl.GRPCSrv.Stop()
	log.SugaredLogger.Info("program was stopped")
}
