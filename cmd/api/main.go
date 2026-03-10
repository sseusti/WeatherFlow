package main

import (
	"WeatherFlow/internal/router"
	"log"
)

func main() {
	r := router.NewRouter()

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
