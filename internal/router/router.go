package router

import (
	"WeatherFlow/internal/handler"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/health", handler.Health)

	api := r.Group("/api/v1")
	api.GET("/health", handler.Health)

	return r
}
