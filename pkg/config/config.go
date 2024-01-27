package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	PGConfig   PGConfig
	GRPCConfig GRPCConfig

	Salt string `env:"SALT"`
}

type PGConfig struct {
	ConnectionString string `env:"PG_CONNECTION_STRING"`
}

type GRPCConfig struct {
	Port int `env:"GRPC_ADDRESS"`
}

var (
	once   sync.Once
	config *Config
)

func GetConfig() (*Config, error) {
	once.Do(func() {
		if err := cleanenv.ReadEnv(&config); err != nil {
			log.Fatal("failed to parse configs: %w", err)
		}
	})

	return config, nil
}
