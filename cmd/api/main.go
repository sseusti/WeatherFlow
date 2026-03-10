package main

import (
	"WeatherFlow/internal/config"
	"WeatherFlow/internal/router"
	"log"
)

func main() {
	cfg := config.Load()
	r := router.New()

	err := r.Run(":" + cfg.Port)
	if err != nil {
		log.Fatal(err)
	}
}
