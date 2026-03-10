package config

import "os"

type Config struct {
	Port              string
	WeatherAPIBaseURL string
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

	return &Config{
		Port: port,
	}
}
