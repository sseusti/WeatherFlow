package service

import "github.com/gin-gonic/gin"

type WeatherService struct{}

func NewWeatherService() *WeatherService {
	return &WeatherService{}
}

func (s *WeatherService) GetCurrent(city string) gin.H {
	return gin.H{
		"city":        city,
		"temperature": 0,
		"condition":   "stub",
	}
}

func (s *WeatherService) GetByCity(city string) gin.H {
	return gin.H{
		"city":        city,
		"source":      "path",
		"temperature": 0,
		"condition":   "stub",
	}
}
