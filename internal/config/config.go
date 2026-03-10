package config

import (
	"os"
)

type Config struct {
	Port              string
	WeatherAPIBaseURL string
	RequestTimeout    string
}

func Load() *Config {
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8080"
	}

	weatherAPIBaseURL := os.Getenv("WEATHER_API_BASE_URL")
	if weatherAPIBaseURL == "" {
		weatherAPIBaseURL = "https://api.open-meteo.com"
	}

	requestTimeout := os.Getenv("REQUEST_TIMEOUT")
	if requestTimeout == "" {
		requestTimeout = "5"
	}

	return &Config{
		Port:              port,
		WeatherAPIBaseURL: weatherAPIBaseURL,
		RequestTimeout:    requestTimeout,
	}
}
