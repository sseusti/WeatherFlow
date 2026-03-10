package client

type WeatherClient struct {
	baseURL string
}

func NewWeatherClient(baseURL string) *WeatherClient {
	return &WeatherClient{
		baseURL: baseURL,
	}
}

func (c *WeatherClient) BaseURL() string {
	return c.baseURL
}
