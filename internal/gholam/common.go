package gholam

import (
	"bytes"
	"io"
	"net/http"

	"go.uber.org/zap"
)

func (c *gholam) sendRequest(url string, token string, payload []byte, method string) ([]byte, error) {
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

	return body, nil
}
