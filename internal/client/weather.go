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
		Timezone    string  `json:"timezone"`
		Elevation   float64 `json:"elevation"`
		Latitude    float64 `json:"latitude"`
		Longitude   float64 `json:"longitude"`
	} `json:"results"`
}

type hourlyForecastResponse struct {
	Hourly struct {
		Time        []string  `json:"time"`
		Temperature []float64 `json:"temperature_2m"`
	} `json:"hourly"`
}

type HourlyForecastPoint struct {
	Time        string
	Temperature float64
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

type GeocodedLocation struct {
	Name        string
	Country     string
	CountryCode string
	Timezone    string
	Elevation   float64
	Latitude    float64
	Longitude   float64
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

func (c *WeatherClient) HourlyForecastURL(lat, lon string) string {
	u, _ := url.Parse(c.baseURL)
	u.Path = "/v1/forecast"

	q := u.Query()
	q.Set("latitude", lat)
	q.Set("longitude", lon)
	q.Set("hourly", "temperature_2m")

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

func (c *WeatherClient) GeocodeCity(ctx context.Context, city string) (GeocodedLocation, error) {
	u := c.GeocodingURL(city)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return GeocodedLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return GeocodedLocation{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return GeocodedLocation{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var geoResp geocodingResponse
	err = json.NewDecoder(resp.Body).Decode(&geoResp)
	if err != nil {
		return GeocodedLocation{}, err
	}

	if len(geoResp.Results) == 0 {
		return GeocodedLocation{}, fmt.Errorf("city not found")
	}

	return GeocodedLocation{geoResp.Results[0].Name,
		geoResp.Results[0].Country,
		geoResp.Results[0].CountryCode,
		geoResp.Results[0].Timezone,
		geoResp.Results[0].Elevation,
		geoResp.Results[0].Latitude,
		geoResp.Results[0].Longitude,
	}, nil
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

func (c *WeatherClient) HourlyForecast(ctx context.Context, lat, lon string) ([]HourlyForecastPoint, error) {
	u := c.HourlyForecastURL(lat, lon)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var hourlyResp hourlyForecastResponse
	err = json.NewDecoder(resp.Body).Decode(&hourlyResp)
	if err != nil {
		return nil, err
	}

	if len(hourlyResp.Hourly.Time) == 0 || len(hourlyResp.Hourly.Temperature) == 0 {
		return nil, fmt.Errorf("forecast not found")
	}
	if len(hourlyResp.Hourly.Time) != len(hourlyResp.Hourly.Temperature) {
		return nil, fmt.Errorf("forecast time length mismatch")
	}

	var hourlyPoints []HourlyForecastPoint
	for i := 0; i < len(hourlyResp.Hourly.Time); i++ {
		hourlyPoints = append(hourlyPoints, HourlyForecastPoint{
			Time:        hourlyResp.Hourly.Time[i],
			Temperature: hourlyResp.Hourly.Temperature[i],
		})
	}

	return hourlyPoints, nil
}
