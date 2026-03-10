package client

import "time"

type WeatherClient struct {
	baseURL string
	timeout time.Duration
}

func NewWeatherClient(baseURL string, timeout time.Duration) *WeatherClient {
	return &WeatherClient{
		baseURL: baseURL,
		timeout: timeout,
	}
}

func (c *WeatherClient) BaseURL() string {
	return c.baseURL
}

func (c *WeatherClient) Timeout() time.Duration {
	return c.timeout
}
