package config

import (
	"github.com/hobs-ai/gholam-cli/internal/gholam"
	"github.com/hobs-ai/gholam-cli/pkg/logger"
)

type Config struct {
	Gholam *gholam.Config `koanf:"gholam"`
	Logger *logger.Config `koanf:"logger"`
}
