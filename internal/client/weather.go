package client

import (
	"context"
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

func (c *WeatherClient) CurrentWeatherStatus(ctx context.Context, city string) (int, error) {
	u := c.CurrentWeatherURL(city)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return 0, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	return resp.StatusCode, nil
}

func (c *WeatherClient) ForecastURL(lat, lon string) string {
	u, _ := url.Parse(c.baseURL)
	u.Path = "/v1/forecast"

	q := u.Query()
	q.Set("latitude", lat)
	q.Set("longitude", lon)
	q.Set("current", "temperature_2m")

	u.RawQuery = q.Encode()

	return u.String()
}

func (c *WeatherClient) GeocodingURL(city string) string {
	u, _ := url.Parse("https://geocoding-api.open-meteo.com")
	u.Path = "/v1/search"

	q := u.Query()
	q.Set("name", city)
	q.Set("count", "1")

	u.RawQuery = q.Encode()

	return u.String()
}
