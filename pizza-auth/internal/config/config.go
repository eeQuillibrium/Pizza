package config

import (
	"flag"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Env         string        `yaml:"env"`
	Storage     Storage       `yaml:"storage"`
	StoragePath string        `yaml:"storage_path"`
	GRPC        GRPCConfig    `yaml:"grpc"`
	TokenTTL    time.Duration `yaml:"token_ttl"`
}
type Storage struct {
	Host     string `yaml:"host"`
	SSLMode  string `yaml:"sslmode"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"dbname"`
	Username string `yaml:"username"`
}
type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func InitConfig() *Config {

	path := fetchConfigPath()
	if path == "" {
		log.Fatal("path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatal("config path is empty:", err)
	}

	var cfg Config

	data, _ := os.ReadFile(path)

	err := yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatal("unmarshal problem occured:", err)
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string
	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
