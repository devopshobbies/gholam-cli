package config

import (
	"github.com/hobs-ai/gholam-cli/pkg/logger"
)

type Config struct {
	Logger *logger.Config `koanf:"logger"`
}
