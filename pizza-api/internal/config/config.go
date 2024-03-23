package config

import (
	"flag"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// grpc, rest, storage
type Config struct {
	Rest Rest `yaml:"rest"`
	GRPC GRPC `yaml:"grpc"`
	Repo Repo `yaml:"repo"`
}

type Rest struct {
	Port int `yaml:"port"`
}

type GRPC struct {
	Auth    Auth    `yaml:"auth"`
	Kitchen Kitchen `yaml:"kitchen"`
	Server  Server  `yaml:"server"`
	Appnum  int     `yaml:"appnum"`
}
type Auth struct {
	Port int `yaml:"port"`
}
type Kitchen struct {
	Port int `yaml:"port"`
}
type Server struct {
	Port int `yaml:"port"`
}

type Repo struct {
	Redis    Redis    `yaml:"redis"`
	Postgres Postgres `yaml:"postgres"`
}
type Redis struct {
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}
type Postgres struct {
	Username string `yaml:"username"`
	Host     string `yaml:"host"`
	DBName   string `yaml:"dbname"`
	Port     int    `yaml:"port"`
	SSLMode  string `yaml:"sslmode"`
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
