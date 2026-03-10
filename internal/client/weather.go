package client

type WeatherClient struct {
	baseURL string
	timeout string
}

func NewWeatherClient(baseURL string, timeout string) *WeatherClient {
	return &WeatherClient{
		baseURL: baseURL,
		timeout: timeout,
	}
}

func (c *WeatherClient) BaseURL() string {
	return c.baseURL
}

func (c *WeatherClient) Timeout() string {
	return c.timeout
}
