package gholam

func (c *gholam) sendRequest(url string, payload []byte, method string) error {
	// req, _ := http.NewRequest(method, url, bytes.NewBuffer(payload))
	// req.Header.Set("Authorization", "Bearer "+c.config.Token)
	// req.Header.Set("Content-Type", "application/json")

	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	c.logger.Error("Error performing request", zap.Error(err))
	// 	return err
	// }
	// defer resp.Body.Close()

	return nil
}
