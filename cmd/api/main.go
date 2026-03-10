package main

import (
	"WeatherFlow/internal/handler"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.GET("/health", handler.Health)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
