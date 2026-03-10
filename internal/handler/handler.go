package handler

import "WeatherFlow/internal/service"

type Handler struct {
	weatherService *service.WeatherService
}

func New() *Handler {
	return &Handler{
		weatherService: service.NewWeatherService(),
	}
}
