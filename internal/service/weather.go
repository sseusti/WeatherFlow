package service

import "WeatherFlow/internal/client"

type WeatherService struct {
	client *client.WeatherClient
}

type CurrentWeatherResponse struct {
	City        string `json:"city"`
	Temperature int    `json:"temperature"`
	Condition   string `json:"condition"`
	SourceURL   string `json:"source_url"`
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

func (s *WeatherService) GetCurrent(city string) CurrentWeatherResponse {
	url := s.client.CurrentWeatherURL(city)
	return CurrentWeatherResponse{
		City:        city,
		Temperature: 0,
		Condition:   "stub",
		SourceURL:   url,
	}
}

func (s *WeatherService) GetByCity(city string) CityWeatherResponse {
	return CityWeatherResponse{
		City:        city,
		Temperature: 0,
		Condition:   "stub",
		Source:      "path",
	}
}
