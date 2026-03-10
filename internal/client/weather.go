package client

import (
	"net/http"
	"time"
)

type WeatherClient struct {
	baseURL string
	timeout time.Duration
	client  *http.Client
}

func NewWeatherClient(baseURL string, timeout time.Duration) *WeatherClient {
	client := &http.Client{Timeout: timeout}
	return &WeatherClient{
		baseURL: baseURL,
		timeout: timeout,
		client:  client,
	}
}

func (c *WeatherClient) BaseURL() string {
	return c.baseURL
}

func (c *WeatherClient) Timeout() time.Duration {
	return c.timeout
}

func (c *WeatherClient) HTTPClient() *http.Client {
	return c.client
}
