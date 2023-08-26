package helper

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"go.uber.org/zap"
)

type Credential struct {
	Token string `json:"token"`
}

func (h *helper) SaveToken(token string) error {
	// Encode the token to base64
	encodedToken := base64.StdEncoding.EncodeToString([]byte(token))

	// Create a configuration struct
	config := Credential{Token: encodedToken}

	configPath, err := h.getPath()
	if err != nil {
		h.logger.Error("Error getting config path", zap.Error(err))
		return err
	}

	// Create the directory if it doesn't exist
	if err = os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
		fmt.Println("Error creating directory:", err)
		return err
	}

	// Open the configuration file for writing
	file, err := os.Create(configPath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close()

	// Encode the configuration struct as JSON and write to the file
	encoder := json.NewEncoder(file)
	if err = encoder.Encode(config); err != nil {
		fmt.Println("Error encoding JSON:", err)
		return err
	}

	return nil
}

func (h *helper) GetToken() (string, error) {
	configPath, err := h.getPath()
	if err != nil {
		h.logger.Error("Error getting config path", zap.Error(err))
		return "", err
	}

	// Open the configuration file for reading
	file, err := os.Open(configPath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return "", err
	}
	defer file.Close()

	// Decode the JSON content from the file into the configuration struct
	var config Credential
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return "", err
	}

	// Decode the base64 encoded token
	decodedToken, err := base64.StdEncoding.DecodeString(config.Token)
	if err != nil {
		fmt.Println("Error decoding token:", err)
		return "", err
	}

	return string(decodedToken), nil
}

func (h *helper) DeleteToken() error {
	configPath, err := h.getPath()
	if err != nil {
		h.logger.Error("Error getting config path", zap.Error(err))
		return err
	}

	if err := os.Remove(configPath); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist.")
			return nil
		}

		return err
	}

	return nil
}

func (h *helper) getPath() (string, error) {
	usr, err := user.Current()
	if err != nil {
		fmt.Println("Error getting user's home directory:", err)
		return "", err
	}

	return filepath.Join(usr.HomeDir, ".config", "gholam", "config.json"), nil

}
