package gholam

import "go.uber.org/zap"

type Gholam interface {
	// Auth related methods
	Register(email string, password string) (string, error)
	Login(email string, password string) (string, error)
	Logout(token string) error
}

type gholam struct {
	config *Config
	logger *zap.Logger
}

func New(cfg *Config, lg *zap.Logger) Gholam {
	return &gholam{config: cfg, logger: lg}
}
