package gholam

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"go.uber.org/zap"
)

type Gholam interface {
	// Auth related methods
	Register(email string, password string) (string, error)
	Login(email string, password string) (string, error)
	Logout(token string) error

	Translate(token, language, phrase string) (string, error)
}

type gholam struct {
	config *Config
	logger *zap.Logger
}

func New(cfg *Config, lg *zap.Logger) Gholam {
	return &gholam{config: cfg, logger: lg}
}

// sendRequest returns the body, status code and error if any error occured
func (c *gholam) sendRequest(url string, token string, payload []byte, method string) (any, error) {
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	if len(token) != 0 {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		c.logger.Error("Error performing request", zap.Error(err))
		return nil, err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		c.logger.Error("Error reading response body", zap.Error(err))
		return nil, err
	}

	type authResponse struct {
		Data    any    `json:"data"`
		Message string `json:"message"`
	}

	var person authResponse
	if err := json.Unmarshal(body, &person); err != nil {
		c.logger.Error("Error unmarshal the response", zap.Error(err))
		return "", err
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		errString := "An error occured during handling the request"
		c.logger.Error(errString, zap.String("Message", person.Message))
		return "", errors.New(errString)
	}

	return person.Data, nil
}

func (instance *gholam) Register(email, password string) (string, error) {
	url := instance.config.Backend + "/auth/register"
	return instance.authenticateRequest(url, email, password)
}

func (instance *gholam) Login(email string, password string) (string, error) {
	url := instance.config.Backend + "/auth/login"
	return instance.authenticateRequest(url, email, password)
}

func (instance *gholam) authenticateRequest(url, email, password string) (string, error) {
	request := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{Email: email, Password: password}

	body, _ := json.Marshal(request)
	responseRaw, err := instance.sendRequest(url, "", body, "POST")
	if err != nil {
		instance.logger.Error("Error authenticating to the Gholam", zap.Error(err))
		return "", err
	}

	responseMap, ok := responseRaw.(map[string]any)
	if !ok {
		errString := "Error reading response"
		instance.logger.Error(errString, zap.Any("response", responseRaw))
		return "", errors.New(errString)
	}

	tokenRaw, exists := responseMap["Token"]
	if !exists {
		errString := "Error missing token"
		instance.logger.Error(errString, zap.Any("response", responseMap))
		return "", errors.New(errString)
	}

	token, ok := tokenRaw.(string)
	if !ok {
		errString := "Error invalid token type"
		instance.logger.Error(errString, zap.Any("response", responseRaw))
		return "", errors.New(errString)
	}

	if len(token) == 0 {
		errString := "Error invalid token length"
		instance.logger.Error(errString, zap.Any("response", responseMap))
		return "", errors.New(errString)
	}

	return token, nil
}

func (instance *gholam) Logout(token string) error {
	url := instance.config.Backend + "/auth/logout"
	if _, err := instance.sendRequest(url, token, []byte{}, "POST"); err != nil {
		instance.logger.Error("Error authenticating to the Gholam", zap.Error(err))
		return err
	}

	return nil
}

func (instance *gholam) Translate(token, language, phrase string) (string, error) {
	request := struct {
		Language string `json:"language"`
		Phrase   string `json:"phrase"`
	}{Language: language, Phrase: phrase}

	body, _ := json.Marshal(request)
	url := instance.config.Backend + "/translate"
	responseRaw, err := instance.sendRequest(url, token, body, "POST")
	if err != nil {
		instance.logger.Error("Error authenticating to the Gholam", zap.Error(err))
		return "", err
	}

	translation, ok := responseRaw.(string)
	if !ok {
		errString := "Error invalid translation type"
		instance.logger.Error(errString, zap.Any("response", responseRaw))
		return "", errors.New(errString)
	}

	if len(translation) == 0 {
		errString := "Error invalid translation length"
		instance.logger.Error(errString, zap.Any("response", responseRaw))
		return "", errors.New(errString)
	}

	return translation, nil
}
