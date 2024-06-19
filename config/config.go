package config

import (
	"errors"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App    `yaml:"app"`
		HTTP   `yaml:"http"`
		Launch `yaml:"launch"`
	}

	App struct {
		Name    string `env-required:"true" yaml:"name"`
		Version string `env-required:"true" yaml:"version"`
	}

	HTTP struct {
		Port string `env-required:"true" yaml:"port"`
	}

	Launch struct {
		Mode string `env-required:"true" yaml:"mode"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)

	if cfg.Launch.Mode != "prod" && cfg.Launch.Mode != "dev" {
		return nil, errors.New("launch mode must be either 'prod' or 'dev'")
	}

	if err != nil {
		return nil, err
	}

	return cfg, nil
}
