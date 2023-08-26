package main

import (
	"log"

	"github.com/hobs-ai/gholam-cli/cmd"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func main() {
	const description = "Gholam CLI is a tool to interact to the Gholam which is an AI-powerd translator application"
	root := &cobra.Command{Short: description}

	root.AddCommand(
		cmd.Register{}.Command(),
		cmd.Login{}.Command(),
		cmd.Logout{}.Command(),
		cmd.Translate{}.Command(),
	)

	if err := root.Execute(); err != nil {
		log.Fatal("failed to execute root command", zap.Error(err))
	}
}
