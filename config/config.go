package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App  `yaml:"app"`
		Psql `yaml:"postgres"`
	}

	App struct {
		Token string `env:"TELEGRAM_TOKEN" env-required:"true"`
		Debug bool   `env:"TELEGRAM_DEBUG" env-required:"true"`
	}

	Psql struct {
		Host         string `env:"DB_HOST" env-default:"localhost"`
		Port         string `env:"DB_PORT" env-default:"5432"`
		User         string `env:"DB_USER" env-required:"true"`
		Password     string `env:"DB_PASSWORD" env-required:"true"`
		DatabaseName string `env:"DB_NAME" env-required:"true"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := new(Config)

	err := cleanenv.ReadConfig(".env", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
