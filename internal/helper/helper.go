package helper

import "go.uber.org/zap"

type Helper interface {
	GetEmailAndPassword() (string, string, error)

	SaveToken(token string) error
	GetToken() (string, error)
	DeleteToken() error
}

type helper struct {
	logger *zap.Logger
}

func New(lg *zap.Logger) Helper {
	return &helper{logger: lg}
}
