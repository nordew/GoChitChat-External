package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	PGConfig   PGConfig
	GRPCConfig GRPCConfig

	Salt string
}

type PGConfig struct {
	ConnectionString string `env:"PG_CONNECTION_STRING"`
}

type GRPCConfig struct {
	Adress string `env:"GRPC_ADRESS"`
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
