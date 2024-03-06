package config

import (
	"flag"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	GRPC GRPCApp `yaml:"grpcapp"`
	REST RESTApp `yaml:"restapp"`
}
type GRPCApp struct {
	Kitchenapi Kitcheapi  `yaml:"kitchenapi"`
	Kitchendel Kitchendel `yaml:"kitchendel"`
	Appnum     int        `yaml:"appnum"`
}
type Kitcheapi struct {
	Port int `yaml:"port"`
}
type Kitchendel struct {
	Port int `yaml:"port"`
}
type RESTApp struct {
	Port int `yaml:"port"`
}

func New() *Config {

	path := fetchConfigPath()

	if path == "" {
		log.Fatal("cfg path is empty")
	}

	if _, err := os.Stat(path); os.IsExist(err) {
		log.Fatalf("file with path %s does not exist: %v", path, err)
	}

	data, _ := os.ReadFile(path)

	var cfg Config

	err := yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatalf("error with cfg unmarshal: %v", err)
	}

	return &cfg
}

func fetchConfigPath() string {
	var path string
	flag.StringVar(&path, "configpath", "", "path for kitchen config")
	flag.Parse()

	if path == "" {
		os.Getenv("CONFIG_PATH")
	}

	return path
}
