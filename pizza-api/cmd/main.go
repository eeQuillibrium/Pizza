package main

import (
	"github.com/eeQuillibrium/pizza-api/internal/app"
	"github.com/eeQuillibrium/pizza-api/internal/logger"
	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

func main() {

	log := logger.New()

	if err := godotenv.Load(); err != nil {
		log.Fatalf(".env reading err: %v", err)
	}

	app := app.New(log)
	app.Run()
}
