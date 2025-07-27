package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env" env-default:"development"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}
type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"0.0.0.0:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"5s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func Mustload() *Config {
	configPath := os.Getenv("./config/local.yaml./your-app")
	if configPath == "" {
		configPath = "config/config.yaml"
		log.Printf("Using default config path: %s", configPath)
	}
	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("error opening config file %s", err)
	}
	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("error opening config file %s", err)
	}
	return &cfg
}
