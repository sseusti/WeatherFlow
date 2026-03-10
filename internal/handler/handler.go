package handler

import (
	"WeatherFlow/internal/client"
	"WeatherFlow/internal/service"
)

type Handler struct {
	weatherService *service.WeatherService
}

func New(client client.WeatherClient) *Handler {
	return &Handler{
		weatherService: service.NewWeatherService(client),
	}
}
