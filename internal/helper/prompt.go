package helper

import (
	"errors"
	"regexp"

	"github.com/manifoldco/promptui"
)

const emailRegexPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`
const passwordRegexPattern = `^(?=.*[a-z])(?=.*[A-Z]).{6,}$`

func (h *helper) GetEmailAndPassword() (string, string, error) {
	email, err := (&promptui.Prompt{
		Label: "Enter your email",
		Validate: func(input string) error {
			if !regexp.MustCompile(emailRegexPattern).MatchString(input) {
				errString := "Invalid email has been given"
				return errors.New(errString)
			}
			return nil
		},
	}).Run()

	if err != nil {
		return "", "", err
	}

	password, err := (&promptui.Prompt{
		Label: "Enter your password",
		Validate: func(input string) error {
			if !regexp.MustCompile(passwordRegexPattern).MatchString(input) {
				errString := "Invalid password has been given"
				return errors.New(errString)
			}
			return nil
		},
	}).Run()

	if err != nil {
		return "", "", err
	}

	return email, password, nil
}
