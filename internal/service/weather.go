package service

import "WeatherFlow/internal/client"

type WeatherService struct {
	client *client.WeatherClient
}

type CurrentWeatherResponse struct {
	City           string `json:"city"`
	Temperature    int    `json:"temperature"`
	Condition      string `json:"condition"`
	SourceURL      string `json:"source_url"`
	ExternalStatus int    `json:"external_status"`
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
	status, err := s.client.CurrentWeatherStatus(city)
	if err != nil {
		status = 0
	}

	return CurrentWeatherResponse{
		City:           city,
		Temperature:    0,
		Condition:      "stub",
		SourceURL:      url,
		ExternalStatus: int(status),
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
