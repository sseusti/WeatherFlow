package service

import (
	"WeatherFlow/internal/client"
	"context"
	"strconv"
)

type WeatherService struct {
	client *client.WeatherClient
}

type CurrentWeatherResponse struct {
	City        string  `json:"city"`
	Temperature float64 `json:"temperature"`
	Condition   string  `json:"condition"`
	SourceURL   string  `json:"source_url"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	FeelsLike   float64 `json:"feels_like"`
	WindSpeed   float64 `json:"wind_speed"`
	Humidity    float64 `json:"humidity"`
}

type CityWeatherResponse struct {
	City        string `json:"city"`
	Temperature int    `json:"temperature"`
	Condition   string `json:"condition"`
	Source      string `json:"source"`
}

func NewWeatherService(client *client.WeatherClient) *WeatherService {
	return &WeatherService{
		client: client,
	}
}

func (s *WeatherService) GetCurrent(ctx context.Context, city string) (CurrentWeatherResponse, error) {
	lat, lon, err := s.client.GeocodeCity(ctx, city)
	if err != nil {
		return CurrentWeatherResponse{}, err
	}

	latStr := strconv.FormatFloat(lat, 'f', -1, 64)
	lonStr := strconv.FormatFloat(lon, 'f', -1, 64)

	url := s.client.ForecastURL(latStr, lonStr)

	forecast, err := s.client.CurrentForecast(ctx, latStr, lonStr)
	if err != nil {
		return CurrentWeatherResponse{}, err
	}

	return CurrentWeatherResponse{
		City:        city,
		Temperature: forecast.Temperature,
		Condition:   mapWeatherCode(forecast.WeatherCode),
		SourceURL:   url,
		Latitude:    lat,
		Longitude:   lon,
		FeelsLike:   forecast.FeelsLike,
		WindSpeed:   forecast.WindSpeed,
		Humidity:    forecast.Humidity,
	}, nil
}

func (s *WeatherService) GetByCity(city string) CityWeatherResponse {
	return CityWeatherResponse{
		City:        city,
		Temperature: 0,
		Condition:   "stub",
		Source:      "path",
	}
}

func mapWeatherCode(code int) string {
	mapper := map[int]string{
		0:  "Clear",
		1:  "Cloudy",
		2:  "Cloudy",
		3:  "Cloudy",
		45: "Fog",
		48: "Fog",
		51: "Drizzle",
		53: "Drizzle",
		55: "Drizzle",
		61: "Rain",
		63: "Rain",
		65: "Rain",
	}

	codeStr, ok := mapper[code]
	if !ok {
		return "Unknown"
	}

	return codeStr
}
