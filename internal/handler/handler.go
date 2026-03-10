package handler

import "WeatherFlow/internal/service"

type Handler struct {
	WeatherService *service.WeatherService
}

func New() *Handler {
	return &Handler{
		WeatherService: service.NewWeatherService(),
	}
}
