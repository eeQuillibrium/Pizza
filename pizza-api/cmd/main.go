package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/eeQuillibrium/pizza-api/internal/app"
	"github.com/eeQuillibrium/pizza-api/internal/handler"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatal("Config reading problem", err)
	}

	handl := handler.New(viper.GetString("grpc.port"))
	
	app := app.New(viper.GetString("server.port"), handl.InitRoutes())

	app.RESTServ.Run("grpc.port")

	//graceful shutdown
}

func initConfig() error {
	viper.AddConfigPath("internal/configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
