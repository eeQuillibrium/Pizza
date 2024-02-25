package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/eeQuillibrium/pizza-auth/internal/app"
	"github.com/eeQuillibrium/pizza-auth/internal/config"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Env var didn't loaded: %v", err)
	}

	cfg := config.InitConfig()

	appl := app.New(cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)

	go appl.GRPCSrv.Run()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop

	log.Print("try to stop the program with ", sign)
	appl.GRPCSrv.Stop()

	log.Print("program was stopped")
}
