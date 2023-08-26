package cmd

import (
	"github.com/hobs-ai/gholam-cli/internal/config"
	"github.com/hobs-ai/gholam-cli/internal/gholam"
	"github.com/hobs-ai/gholam-cli/internal/helper"
	"github.com/hobs-ai/gholam-cli/pkg/logger"
	"go.uber.org/zap"

	"github.com/spf13/cobra"
)

type Logout struct{}

func (cmd Logout) Command() *cobra.Command {
	run := func(_ *cobra.Command, _ []string) {
		cmd.main(config.Load(true))
	}

	return &cobra.Command{
		Use:   "logout",
		Short: "Logout from the gholam backend",
		Run:   run,
	}
}

func (cmd *Logout) main(cfg *config.Config) {
	logger := logger.NewZap(cfg.Logger)

	token, err := helper.GetToken()
	if err != nil {
		logger.Error("Failed to get the token", zap.Error(err))
		return
	}

	gholam := gholam.New()

	if err := gholam.Logout(token); err != nil {
		logger.Error("Failed to logout from the Gholam", zap.Error(err))
		return
	}

	logger.Info("You are now logged out from the Gholam")
}
