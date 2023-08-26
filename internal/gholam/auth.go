package gholam

import (
	"encoding/json"

	"go.uber.org/zap"
)

type auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (instance *gholam) Register(email, password string) (string, error) {
	url := instance.config.Backend + "/auth/register"
	body, _ := json.Marshal(auth{Email: email, Password: password})
	response, err := instance.sendRequest(url, "", body, "POST")
	if err != nil {
		instance.logger.Error("Error registring to the Gholam", zap.Error(err))
		return "", err
	}

	instance.logger.Info(string(response))

	return string(response), nil
}

func (instance *gholam) Login(email string, password string) (string, error) {
	return "", nil
}

func (instance *gholam) Logout(token string) error {
	return nil
}
