package cmd

import (
	"os"

	"github.com/hobs-ai/gholam-cli/internal/config"
	"github.com/hobs-ai/gholam-cli/pkg/logger"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type Translate struct{}

func (cmd Translate) Command(trap chan os.Signal) *cobra.Command {
	run := func(_ *cobra.Command, _ []string) {
		cmd.main(config.Load(true), trap)
	}

	return &cobra.Command{
		Use:   "translate",
		Short: "Translate the given sentences",
		Run:   run,
	}
}

func (cmd *Translate) main(cfg *config.Config, trap chan os.Signal) {
	logger := logger.NewZap(cfg.Logger)

	// Keep this at the bottom of the main function
	field := zap.String("signal trap", (<-trap).String())
	logger.Info("exiting by receiving a unix signal", field)
}
