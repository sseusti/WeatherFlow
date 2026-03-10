package app

import (
	"WeatherFlow/internal/client"
	"WeatherFlow/internal/config"
	"WeatherFlow/internal/handler"
	"WeatherFlow/internal/router"

	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
}

func New(cfg *config.Config) *App {
	weatherClient := client.NewWeatherClient(cfg.WeatherAPIBaseURL, cfg.RequestTimeout)
	h := handler.New(weatherClient)
	r := router.New(h)

	return &App{
		Router: r,
	}
}
