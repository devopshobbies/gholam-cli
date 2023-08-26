package config

import (
	"github.com/hobs-ai/gholam-cli/internal/gholam"
	"github.com/hobs-ai/gholam-cli/pkg/logger"
)

func defaultConfig() *Config {
	return &Config{
		Gholam: &gholam.Config{
			Backend: "",
		},
		Logger: &logger.Config{
			Development: true,
			Level:       "info",
			Encoding:    "console",
		},
	}
}
