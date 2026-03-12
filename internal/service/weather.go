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
	City           string  `json:"city"`
	Temperature    float64 `json:"temperature"`
	Condition      string  `json:"condition"`
	SourceURL      string  `json:"source_url"`
	ExternalStatus int     `json:"external_status"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
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

	status, err := s.client.ForecastStatus(ctx, latStr, lonStr)
	if err != nil {
		return CurrentWeatherResponse{}, err
	}

	url := s.client.ForecastURL(latStr, lonStr)

	temp, err := s.client.CurrentTemperature(ctx, latStr, lonStr)
	if err != nil {
		return CurrentWeatherResponse{}, err
	}

	return CurrentWeatherResponse{
		City:           city,
		Temperature:    temp,
		Condition:      "stub",
		SourceURL:      url,
		ExternalStatus: status,
		Latitude:       lat,
		Longitude:      lon,
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
