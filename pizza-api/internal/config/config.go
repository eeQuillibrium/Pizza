package config

import (
	"flag"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// grpc, rest, storage
type Config struct {
	Server Server
	GRPC   GRPC
}
type Server struct {
	Port string `yaml:"port"`
}
type GRPC struct {
	Port   string `yaml:"port"`
	Appnum int    `yaml:"appnum"`
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

	flag.StringVar(&path, "config", "", "path to config")
	flag.Parse()

	if path == "" {
		path = os.Getenv("CONFIG_PATH")
	}
	return path
}
