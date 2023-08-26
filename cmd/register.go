package cmd

import (
	"github.com/hobs-ai/gholam-cli/internal/config"
	"github.com/hobs-ai/gholam-cli/internal/gholam"
	"github.com/hobs-ai/gholam-cli/internal/helper"
	"github.com/hobs-ai/gholam-cli/pkg/logger"
	"go.uber.org/zap"

	"github.com/spf13/cobra"
)

type Register struct{}

func (cmd Register) Command() *cobra.Command {
	run := func(_ *cobra.Command, _ []string) {
		cmd.main(config.Load())
	}

	return &cobra.Command{
		Use:   "register",
		Short: "Register to the gholam backend",
		Run:   run,
	}
}

func (cmd *Register) main(cfg *config.Config) {
	logger := logger.NewZap(cfg.Logger)

	helper := helper.New(logger)

	email, password, err := helper.GetEmailAndPassword()
	if err != nil {
		logger.Error("Prompt failed", zap.Error(err))
		return
	}

	gholam := gholam.New(cfg.Gholam, logger)
	token, err := gholam.Register(email, password)
	if err != nil {
		logger.Error("Failed to register to the gholam", zap.Error(err))
		return
	}

	if err := helper.SaveToken(token); err != nil {
		logger.Error("Failed to save the token", zap.Error(err))
		return
	}

	logger.Info("You are now logged in to the Gholam")
}
