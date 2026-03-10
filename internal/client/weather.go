package client

import (
	"net/http"
	"net/url"
	"time"
)

type WeatherClient struct {
	baseURL    string
	timeout    time.Duration
	httpClient *http.Client
}

func NewWeatherClient(baseURL string, timeout time.Duration) *WeatherClient {
	client := &http.Client{Timeout: timeout}
	return &WeatherClient{
		baseURL:    baseURL,
		timeout:    timeout,
		httpClient: client,
	}
}

func (c *WeatherClient) BaseURL() string {
	return c.baseURL
}

func (c *WeatherClient) Timeout() time.Duration {
	return c.timeout
}

func (c *WeatherClient) HTTPClient() *http.Client {
	return c.httpClient
}

func (c *WeatherClient) CurrentWeatherURL(city string) string {
	u, _ := url.Parse(c.baseURL)
	u.Path = "/current"
	q := u.Query()
	q.Set("city", city)
	u.RawQuery = q.Encode()

	return u.String()
}

func (c *WeatherClient) CurrentWeatherStatus(city string) (int, error) {
	u := c.CurrentWeatherURL(city)

	resp, err := c.httpClient.Get(u)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	return resp.StatusCode, nil
}
