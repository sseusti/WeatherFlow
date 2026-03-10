package router

import (
	"WeatherFlow/internal/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger(), gin.Recovery())

	router.GET("/health", handler.Health)

	return router
}
