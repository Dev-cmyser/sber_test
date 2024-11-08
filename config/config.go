// Package config s
package config

import (
	"fmt"
	"path"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config s.
	Config struct {
		App   `yaml:"app"`
		HTTP  `yaml:"http"`
		Log   `yaml:"log"`
		Cache `yaml:"cache"`
	}

	// App s.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// HTTP s.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// Log s.
	Log struct {
		Level string `env-required:"true" yaml:"level" env:"LOG_LEVEL"`
	}

	// Cache s.
	Cache struct {
		TTL  int `env-required:"true" yaml:"ttl_seconds" env:"TTL"`
		SIZE int `env-required:"true" yaml:"size" env:"SIZE"`
	}
)

// NewConfig s.
func NewConfig(configPath string) (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(path.Join("./", configPath), cfg)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	err = cleanenv.UpdateEnv(cfg)
	if err != nil {
		return nil, fmt.Errorf("error updating env: %w", err)
	}

	return cfg, nil
}
