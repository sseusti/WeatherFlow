package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type WeatherClient struct {
	baseURL    string
	timeout    time.Duration
	httpClient *http.Client
}

type geocodingResponse struct {
	Results []struct {
		Name        string  `json:"name"`
		Country     string  `json:"country"`
		CountryCode string  `json:"country_code"`
		Latitude    float64 `json:"latitude"`
		Longitude   float64 `json:"longitude"`
	} `json:"results"`
}

type forecastResponse struct {
	Current struct {
		Temperature         float64 `json:"temperature_2m"`
		WeatherCode         int     `json:"weather_code"`
		ApparentTemperature float64 `json:"apparent_temperature"`
		WindSpeed           float64 `json:"wind_speed_10m"`
		Humidity            float64 `json:"relative_humidity_2m"`
		IsDay               int     `json:"is_day"`
		Precipitation       float64 `json:"precipitation"`
		Time                string  `json:"time"`
	} `json:"current"`
}

type CurrentForecastData struct {
	Temperature   float64
	WeatherCode   int
	FeelsLike     float64
	WindSpeed     float64
	Humidity      float64
	IsDay         int
	Precipitation float64
	Time          string
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

	return c.getStatus(ctx, u)
}

func (c *WeatherClient) ForecastURL(lat, lon string) string {
	u, _ := url.Parse(c.baseURL)
	u.Path = "/v1/forecast"

	q := u.Query()
	q.Set("latitude", lat)
	q.Set("longitude", lon)
	q.Set("current", "temperature_2m,apparent_temperature,weather_code,wind_speed_10m,relative_humidity_2m,is_day,precipitation")

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

func (c *WeatherClient) GeocodingStatus(ctx context.Context, city string) (int, error) {
	u := c.GeocodingURL(city)

	return c.getStatus(ctx, u)
}

func (c *WeatherClient) getStatus(ctx context.Context, rawURL string) (int, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, rawURL, nil)
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

func (c *WeatherClient) GeocodeCity(ctx context.Context, city string) (string, string, string, float64, float64, error) {
	u := c.GeocodingURL(city)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return "", "", "", 0, 0, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", "", "", 0, 0, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", "", 0, 0, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var geoResp geocodingResponse
	err = json.NewDecoder(resp.Body).Decode(&geoResp)
	if err != nil {
		return "", "", "", 0, 0, err
	}

	if len(geoResp.Results) == 0 {
		return "", "", "", 0, 0, fmt.Errorf("city not found")
	}

	return geoResp.Results[0].Name, geoResp.Results[0].Country, geoResp.Results[0].CountryCode, geoResp.Results[0].Latitude, geoResp.Results[0].Longitude, nil
}

func (c *WeatherClient) ForecastStatus(ctx context.Context, lat, lon string) (int, error) {
	u := c.ForecastURL(lat, lon)

	return c.getStatus(ctx, u)
}

func (c *WeatherClient) CurrentForecast(ctx context.Context, lat, lon string) (CurrentForecastData, error) {
	u := c.ForecastURL(lat, lon)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return CurrentForecastData{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return CurrentForecastData{}, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return CurrentForecastData{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var forResp forecastResponse
	err = json.NewDecoder(resp.Body).Decode(&forResp)
	if err != nil {
		return CurrentForecastData{}, err
	}

	return CurrentForecastData{
		Temperature:   forResp.Current.Temperature,
		WeatherCode:   forResp.Current.WeatherCode,
		FeelsLike:     forResp.Current.ApparentTemperature,
		WindSpeed:     forResp.Current.WindSpeed,
		Humidity:      forResp.Current.Humidity,
		IsDay:         forResp.Current.IsDay,
		Precipitation: forResp.Current.Precipitation,
		Time:          forResp.Current.Time,
	}, nil
}
