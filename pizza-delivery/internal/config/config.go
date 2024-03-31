package config

import (
	"flag"
	"os"

	"github.com/eeQuillibrium/pizza-delivery/internal/logger"
	"gopkg.in/yaml.v3"
)

type Config struct {
	REST REST `yaml:"rest"`
	GRPC GRPC `yaml:"grpc"`
	Repo Repo `yaml:"repo"`
}
type REST struct {
	Port int `yaml:"port"`
}
type GRPC struct {
	KitchenClient KitchenClient `yaml:"client"`
	GatewayClient GatewayClient
	Server        Server `yaml:"server"`
}
type KitchenClient struct {
	Port int `yaml:"port"`
}
type GatewayClient struct {
	Port int `yaml:"port"`
}
type Server struct {
	Port int `yaml:"port"`
}
type Repo struct {
	Redis Redis `yaml:"redis"`
}
type Redis struct {
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

func New(log *logger.Logger) *Config {
	path := fetchConfigPath()
	if path == "" {
		log.SugaredLogger.Fatal("empty path")
	}

	var cfg Config

	content, err := os.ReadFile(path)
	if err != nil {
		log.SugaredLogger.Fatalf("file reading error: %w", err)
	}

	yaml.Unmarshal(content, &cfg)
	return &cfg
}
func fetchConfigPath() string {
	var path string

	flag.StringVar(&path, "configpath", "", "path for config")
	flag.Parse()

	if path == "" {
		path = os.Getenv("CONFIG_PATH")
	}

	return path
}
