package main

import (
	"github.com/eeQuillibrium/pizza-api/internal/app"
	"github.com/eeQuillibrium/pizza-api/internal/config"
	"github.com/eeQuillibrium/pizza-api/internal/handler"
	"github.com/spf13/viper"
)

func main() {

	cfg := config.New()

	handl := handler.New(cfg.GRPC.Port)

	app := app.New(cfg.Server.Port, handl.InitRoutes())

	app.RESTServ.Run("grpc.port")

	//graceful shutdown
}

func initConfig() error {
	viper.AddConfigPath("internal/configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
