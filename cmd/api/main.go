package main

import (
	"WeatherFlow/internal/app"
	"WeatherFlow/internal/config"
	"log"
)

func main() {
	cfg := config.Load()
	application := app.New(cfg)

	err := application.Router.Run(":" + cfg.Port)
	if err != nil {
		log.Fatal(err)
	}
}
