package main

import (
	"WeatherFlow/internal/app"
	"WeatherFlow/internal/config"
	"log"
)

func main() {
	cfg := config.Load()
	a := app.New(cfg)

	err := a.Router.Run(":" + cfg.Port)
	if err != nil {
		log.Fatal(err)
	}
}
