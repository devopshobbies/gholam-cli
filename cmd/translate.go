package cmd

import (
	"errors"

	"github.com/hobs-ai/gholam-cli/internal/config"
	"github.com/hobs-ai/gholam-cli/internal/gholam"
	"github.com/hobs-ai/gholam-cli/internal/helper"
	"github.com/hobs-ai/gholam-cli/internal/models"
	"github.com/hobs-ai/gholam-cli/pkg/logger"
	"github.com/manifoldco/promptui"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type Translate struct{}

func (cmd Translate) Command() *cobra.Command {
	run := func(_ *cobra.Command, _ []string) {
		cmd.main(config.Load())
	}

	return &cobra.Command{
		Use:   "translate",
		Short: "Translate the given sentences",
		Run:   run,
	}
}

func (cmd *Translate) main(cfg *config.Config) {
	logger := logger.NewZap(cfg.Logger)

	helper := helper.New(logger)

	token, err := helper.GetToken()
	if err != nil {
		logger.Error("Failed to get the token", zap.Error(err))
		return
	}

	gholam := gholam.New(cfg.Gholam, logger)

	for {
		_, languageRaw, err := (&promptui.Select{
			Label: "Translate the phrase into which language",
			Items: []string{string(models.English), string(models.Persian)},
		}).Run()

		if err != nil {
			logger.Error("Failed to retrieve language", zap.Error(err))
			return
		}

		phrase, err := (&promptui.Prompt{
			Label: "Enter your phrase",
			Validate: func(input string) error {
				if len(input) <= 0 || len(input) > 500 {
					errString := "Invalid phrase length has been given"
					return errors.New(errString)
				}
				return nil
			},
		}).Run()

		if err != nil {
			logger.Error("Failed to retrieve phrase", zap.Error(err))
			return
		}

		languageCode := models.LanguageCodeFromLanguage(models.Language(languageRaw))
		translation, err := gholam.Translate(token, string(languageCode), phrase)
		if err != nil {
			logger.Error("Failed to translate the phrase", zap.Error(err))
			return
		}

		logger.Info(translation)

		_, continueTranslating, err := (&promptui.Select{
			Label: "Continue translating other phrases",
			Items: []string{"Yes", "No"},
		}).Run()

		if err != nil {
			logger.Error("Failed to retrieve confirmation", zap.Error(err))
			return
		} else if continueTranslating == "No" {
			return
		}
	}
}
