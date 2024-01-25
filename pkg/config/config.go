package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	PGConfig PGConfig
}

type PGConfig struct {
	Host     string `env:"PG_HOST"`
	Port     int    `env:"PG_PORT"`
	User     string `env:"PG_USER"`
	Password string `env:"PG_PASSWORD"`
	DBName   string `env:"PG_DB_NAME"`
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
