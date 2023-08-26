package config

import (
	"github.com/hobs-ai/gholam-cli/internal/gholam"
	"github.com/hobs-ai/gholam-cli/pkg/logger"
)

func defaultConfig() *Config {
	return &Config{
		Gholam: &gholam.Config{
			Backend: "https://gholam.mohammadne.me/api/v1",
		},
		Logger: &logger.Config{
			Development: true,
			Level:       "info",
			Encoding:    "console",
		},
	}
}
