package main

import (
	"github.com/eeQuillibrium/pizza-kitchen/internal/app"
	"github.com/eeQuillibrium/pizza-kitchen/internal/config"
)

func main() {

	cfg := config.New()
	
	app := app.New()

	app.Run()
}
