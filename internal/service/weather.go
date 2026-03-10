package service

import (
	"WeatherFlow/internal/client"
	"context"
)

type WeatherService struct {
	client *client.WeatherClient
}

type CurrentWeatherResponse struct {
	City            string `json:"city"`
	Temperature     int    `json:"temperature"`
	Condition       string `json:"condition"`
	SourceURL       string `json:"source_url"`
	ExternalStatus  int    `json:"external_status"`
	GeocodingStatus int    `json:"geocoding_status"`
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
	url := s.client.CurrentWeatherURL(city)
	status, err := s.client.CurrentWeatherStatus(ctx, city)
	if err != nil {
		return CurrentWeatherResponse{}, err
	}

	geocodingStatus, err := s.client.GeocodingStatus(ctx, city)
	if err != nil {
		return CurrentWeatherResponse{}, err
	}

	return CurrentWeatherResponse{
		City:            city,
		Temperature:     0,
		Condition:       "stub",
		SourceURL:       url,
		ExternalStatus:  status,
		GeocodingStatus: geocodingStatus,
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
