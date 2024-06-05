package daily_hot_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	BaseUrl string
}

func New(baseUrl string) DailyHotAPI {
	return &Client{BaseUrl: baseUrl}
}

func (c *Client) fetchHots(endpoint string, dest any) error {
	url := fmt.Sprintf("%s/%s", c.BaseUrl, endpoint)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch data: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, dest)
}
