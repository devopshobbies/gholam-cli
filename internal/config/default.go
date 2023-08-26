package config

import (
	"github.com/hobs-ai/gholam-cli/pkg/logger"
)

func defaultConfig() *Config {
	return &Config{
		Logger: &logger.Config{
			Development: true,
			Level:       "info",
			Encoding:    "console",
		},
	}
}
