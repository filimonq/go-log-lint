package config

import (
	"os"

	"github.com/filimonq/go-log-lint/internal/domain"
	"gopkg.in/yaml.v3"
)

func Load(path string) (*domain.Config, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return domain.DefaultConfig(), nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg := domain.DefaultConfig()
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
