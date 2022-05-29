package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		Log  `yaml:"logger"`
		PG   `yaml:"postgres"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
		GinMode string `env-required:"true" yaml:"gin-mode" env:"GIN_MODE"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// Log -.
	Log struct {
		Env   string `env-required:"true" yaml:"log_env"   env:"APP_ENV"`
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	// PG -.
	PG struct {
		PoolMax  int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
		Host     string `env-required:"true"                 env:"PG_HOST"`
		Port     string `env-required:"true" yaml:"port" env:"PG_PORT"`
		User     string `env-required:"true" yaml:"user" env:"PG_USER"`
		Password string `env-required:"true" yaml:"password" env:"PG_PASSWORD"`
		DBName   string `env-required:"true" yaml:"db_name" env:"PG_DB_NAME"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
