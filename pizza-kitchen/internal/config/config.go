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
	Repo Repo    `yaml:"repo"`
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

type Repo struct {
	Redis Redis `yaml:"redis"`
}
type Redis struct {
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

func New() *Config {

	log.Print("try to load config...")

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

	log.Print("config loaded successful")

	return &cfg
}

func fetchConfigPath() string {
	var path string
	flag.StringVar(&path, "configpath", "", "path for kitchen config")
	flag.Parse()

	if path == "" {
		path = os.Getenv("CONFIG_PATH")
	}

	return path
}
