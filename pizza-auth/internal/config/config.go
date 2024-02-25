package config

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Env         string        `yaml:"env"`
	StoragePath string        `yaml:"storage_path"`
	GRPC        GRPCConfig    `yaml:"grpc"`
	TokenTTL    time.Duration `yaml:"token_ttl"`
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

	data, _ := ioutil.ReadFile(path)

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
