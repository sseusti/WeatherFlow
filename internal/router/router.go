package router

import (
	"WeatherFlow/internal/handler"

	"github.com/gin-gonic/gin"
)

func New(h *handler.Handler) *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/health", h.Health)

	apiV1 := r.Group("/api/v1")
	apiV1.GET("/health", h.Health)

	weather := apiV1.Group("/weather")
	weather.GET("/current", h.GetCurrentWeather)
	weather.GET("/cities/:city", h.GetWeatherByCity)
	weather.GET("/hourly", h.GetHourlyWeatherByCity)

	return r
}
