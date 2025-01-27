package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		Log  `yaml:"log"`
		PG   `yaml:"postgres"`
		Auth `yaml:"auth"`
		JWT  `yaml:"jwt"`
	}

	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	Log struct {
		Level string `env-required:"true" yaml:"log_level" env:"LOG_LEVEL"`
	}

	PG struct {
		PoolMax int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
		URL     string `env-required:"true"  env:"PG_URL"`
	}

	Auth struct {
		GoogleClientID     string `env-required:"true"  env:"GOOGLE_CLIENT_ID"`
		GoogleClientSecret string `env-required:"true"  env:"GOOGLE_CLIENT_SECRET"`
		RedirectURL        string `env-required:"true" yaml:"redirect_url" env:"GOOGLE_REDIRECT_URL"`
	}
	JWT struct {
		Secret string `env-required:"true" env:"JWT_SECRET"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadConfig("./config/config.yml", cfg)

	if err != nil {
		return nil, fmt.Errorf("failed to read configuration: %v", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil

}
