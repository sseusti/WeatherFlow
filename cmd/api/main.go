package main

import (
	"WeatherFlow/internal/router"
	"log"
)

func main() {
	r := router.New()

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
