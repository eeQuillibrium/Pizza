package config

import (
	"flag"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// grpc, rest, storage
type Config struct {
	Server Server `yaml:"server"`
	GRPC   GRPC   `yaml:"grpc"`
	Repo   Repo   `yaml:"repo"`
}

type Server struct {
	Port string `yaml:"port"`
}
type GRPC struct {
	Auth    Auth    `yaml:"auth"`
	Kitchen Kitchen `yaml:"kitchen"`
	Appnum  int     `yaml:"appnum"`
}
type Auth struct {
	Port int `yaml:"port"`
}
type Kitchen struct {
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
	path := fetchConfigPath()

	if path == "" {
		log.Print("config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatal("don't have this file in path:", err)
	}

	data, _ := os.ReadFile(path)

	var cfg Config

	err := yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatal("unmarshal problem occured:", err)
	}

	return &cfg
}
func fetchConfigPath() string {
	var path string

	flag.StringVar(&path, "configpath", "", "path to config")
	flag.Parse()

	if path == "" {
		path = os.Getenv("CONFIG_PATH")
	}
	return path
}
