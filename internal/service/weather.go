package service

import "WeatherFlow/internal/client"

type WeatherService struct {
	Client *client.WeatherClient
}

type CurrentWeatherResponse struct {
	City        string
	Temperature int
	Condition   string
}

type CityWeatherResponse struct {
	City        string
	Temperature int
	Condition   string
	Source      string
}

func NewWeatherService(client *client.WeatherClient) *WeatherService {
	return &WeatherService{
		Client: client,
	}
}

func (s *WeatherService) GetCurrent(city string) CurrentWeatherResponse {
	return CurrentWeatherResponse{
		City:        city,
		Temperature: 0,
		Condition:   "stub",
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
