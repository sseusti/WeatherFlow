package router

import (
	"WeatherFlow/internal/handler"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/health", handler.Health)

	apiV1 := r.Group("/api/v1")
	apiV1.GET("/health", handler.Health)
	apiV1.GET("/weather/current", handler.GetCurrentWeather)
	return r
}
