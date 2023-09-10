package config

import (
	"github.com/hobs-ai/gholam-cli/internal/gholam"
	"github.com/hobs-ai/gholam-cli/pkg/logger"
)

func defaultConfig() *Config {
	return &Config{
		Gholam: &gholam.Config{
			Backend: "http://ns320972.ip-37-187-148.eu:30586/api/v1",
		},
		Logger: &logger.Config{
			Development: true,
			Level:       "info",
			Encoding:    "console",
		},
	}
}
