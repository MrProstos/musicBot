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
		Token string `yaml:"token" env-required:"true"`
		Debug bool   `yaml:"debug" env-required:"true"`
	}

	Psql struct {
		Host         string `yaml:"host" env-default:"localhost"`
		Port         string `yaml:"port" env-default:"5432"`
		User         string `yaml:"user" env-required:"true"`
		Password     string `yaml:"password" env-required:"true"`
		DatabaseName string `yaml:"database_name" env-required:"true"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := new(Config)

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
