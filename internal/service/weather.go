package service

import (
	"WeatherFlow/internal/client"
	"context"
	"strconv"
)

type WeatherService struct {
	client *client.WeatherClient
}

type CurrentWeatherResponse struct {
	SourceURL string                 `json:"source_url"`
	Units     CurrentWeatherUnits    `json:"units"`
	Location  CurrentWeatherLocation `json:"location"`
	Weather   CurrentWeatherData     `json:"weather"`
}

type CurrentWeatherLocation struct {
	City            string                      `json:"city"`
	CityDisplayName string                      `json:"city_display_name"`
	Country         string                      `json:"country"`
	CountryCode     string                      `json:"country_code"`
	Timezone        string                      `json:"timezone"`
	Elevation       float64                     `json:"elevation"`
	Latitude        float64                     `json:"latitude"`
	Longitude       float64                     `json:"longitude"`
	Units           CurrentWeatherLocationUnits `json:"units"`
}

type CurrentWeatherUnits struct {
	Temperature   string `json:"temperature"`
	FeelsLike     string `json:"feels_like"`
	WindSpeed     string `json:"wind_speed"`
	Humidity      string `json:"humidity"`
	Precipitation string `json:"precipitation"`
}

type CityWeatherResponse struct {
	City        string `json:"city"`
	Temperature int    `json:"temperature"`
	Condition   string `json:"condition"`
	Source      string `json:"source"`
}

type CurrentWeatherLocationUnits struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Elevation string `json:"elevation"`
}

type CurrentWeatherData struct {
	Time          string  `json:"time"`
	Temperature   float64 `json:"temperature"`
	FeelsLike     float64 `json:"feels_like"`
	Condition     string  `json:"condition"`
	WindSpeed     float64 `json:"wind_speed"`
	Humidity      float64 `json:"humidity"`
	Precipitation float64 `json:"precipitation"`
	IsDay         bool    `json:"is_day"`
}

type HourlyWeatherPoint struct {
	Time          string  `json:"time"`
	Temperature   float64 `json:"temperature"`
	FeelsLike     float64 `json:"feels_like"`
	Precipitation float64 `json:"precipitation"`
	Condition     string  `json:"condition"`
	IsDay         bool    `json:"is_day"`
	WindSpeed     float64 `json:"wind_speed"`
	Humidity      float64 `json:"humidity"`
}

type HourlyWeatherResponse struct {
	Location CurrentWeatherLocation `json:"location"`
	Hourly   []HourlyWeatherPoint   `json:"hourly"`
	Units    HourlyWeatherUnits     `json:"units"`
}

type HourlyWeatherUnits struct {
	Temperature   string `json:"temperature"`
	FeelsLike     string `json:"feels_like"`
	Precipitation string `json:"precipitation"`
	WindSpeed     string `json:"wind_speed"`
	Humidity      string `json:"humidity"`
}

type DailyWeatherResponse struct {
	Location CurrentWeatherLocation `json:"location"`
	Daily    []DailyWeatherPoint    `json:"daily"`
	Units    DailyWeatherUnits      `json:"units"`
}

type DailyWeatherUnits struct {
	MaxTemperature string `json:"max_temperature"`
	MinTemperature string `json:"min_temperature"`
}

type DailyWeatherPoint struct {
	Date             string  `json:"date"`
	MaxTemperature   float64 `json:"max_temperature"`
	MinTemperature   float64 `json:"min_temperature"`
	WeatherCondition string  `json:"weather_condition"`
}

func NewWeatherService(client *client.WeatherClient) *WeatherService {
	return &WeatherService{
		client: client,
	}
}

func (s *WeatherService) GetCurrent(ctx context.Context, city string) (CurrentWeatherResponse, error) {
	location, err := s.client.GeocodeCity(ctx, city)
	if err != nil {
		return CurrentWeatherResponse{}, err
	}

	latStr := strconv.FormatFloat(location.Latitude, 'f', -1, 64)
	lonStr := strconv.FormatFloat(location.Longitude, 'f', -1, 64)

	url := s.client.ForecastURL(latStr, lonStr)

	forecast, err := s.client.CurrentForecast(ctx, latStr, lonStr)
	if err != nil {
		return CurrentWeatherResponse{}, err
	}

	return CurrentWeatherResponse{
		SourceURL: url,
		Location: CurrentWeatherLocation{
			City:            city,
			CityDisplayName: location.Name,
			Country:         location.Country,
			CountryCode:     location.CountryCode,
			Timezone:        location.Timezone,
			Elevation:       location.Elevation,
			Latitude:        location.Latitude,
			Longitude:       location.Longitude,
			Units: CurrentWeatherLocationUnits{
				Latitude:  "°",
				Longitude: "°",
				Elevation: "m",
			},
		},
		Units: CurrentWeatherUnits{
			Temperature:   "°C",
			FeelsLike:     "°C",
			WindSpeed:     "km/h",
			Humidity:      "%",
			Precipitation: "mm",
		},
		Weather: CurrentWeatherData{
			Time:          forecast.Time,
			Temperature:   forecast.Temperature,
			FeelsLike:     forecast.FeelsLike,
			WindSpeed:     forecast.WindSpeed,
			Humidity:      forecast.Humidity,
			Precipitation: forecast.Precipitation,
			IsDay:         forecast.IsDay == 1,
			Condition:     mapWeatherCode(forecast.WeatherCode),
		},
	}, nil
}

func (s *WeatherService) GetHourly(ctx context.Context, city string) (HourlyWeatherResponse, error) {
	location, err := s.client.GeocodeCity(ctx, city)
	if err != nil {
		return HourlyWeatherResponse{}, err
	}

	latStr := strconv.FormatFloat(location.Latitude, 'f', -1, 64)
	lonStr := strconv.FormatFloat(location.Longitude, 'f', -1, 64)

	points, err := s.client.HourlyForecast(ctx, latStr, lonStr)
	if err != nil {
		return HourlyWeatherResponse{}, err
	}

	locationResp := CurrentWeatherLocation{
		City:            city,
		CityDisplayName: location.Name,
		Country:         location.Country,
		CountryCode:     location.CountryCode,
		Timezone:        location.Timezone,
		Elevation:       location.Elevation,
		Latitude:        location.Latitude,
		Longitude:       location.Longitude,
		Units: CurrentWeatherLocationUnits{
			Latitude:  "°",
			Longitude: "°",
			Elevation: "m",
		},
	}

	hourlyResp := make([]HourlyWeatherPoint, len(points))
	for i := 0; i < len(points); i++ {
		hourlyResp[i] = HourlyWeatherPoint{
			Time:          points[i].Time,
			Temperature:   points[i].Temperature,
			FeelsLike:     points[i].FeelsLike,
			Precipitation: points[i].Precipitation,
			Condition:     mapWeatherCode(points[i].WeatherCode),
			IsDay:         points[i].IsDay == 1,
			WindSpeed:     points[i].WindSpeed,
			Humidity:      points[i].Humidity,
		}
	}

	return HourlyWeatherResponse{
		Location: locationResp,
		Hourly:   hourlyResp,
		Units: HourlyWeatherUnits{
			Temperature:   "°C",
			FeelsLike:     "°C",
			Precipitation: "mm",
			WindSpeed:     "km/h",
			Humidity:      "%",
		},
	}, nil
}

func (s *WeatherService) GetDaily(ctx context.Context, city string) (DailyWeatherResponse, error) {
	location, err := s.client.GeocodeCity(ctx, city)
	if err != nil {
		return DailyWeatherResponse{}, err
	}

	latStr := strconv.FormatFloat(location.Latitude, 'f', -1, 64)
	lonStr := strconv.FormatFloat(location.Longitude, 'f', -1, 64)

	points, err := s.client.DailyForecast(ctx, latStr, lonStr)
	if err != nil {
		return DailyWeatherResponse{}, err
	}

	locationResp := CurrentWeatherLocation{
		City:            city,
		CityDisplayName: location.Name,
		Country:         location.Country,
		CountryCode:     location.CountryCode,
		Timezone:        location.Timezone,
		Elevation:       location.Elevation,
		Latitude:        location.Latitude,
		Longitude:       location.Longitude,
		Units: CurrentWeatherLocationUnits{
			Latitude:  "°",
			Longitude: "°",
			Elevation: "m",
		},
	}

	dailyResp := make([]DailyWeatherPoint, len(points))
	for i := 0; i < len(points); i++ {
		dailyResp[i] = DailyWeatherPoint{
			Date:             points[i].Date,
			MaxTemperature:   points[i].MaxTemperature,
			MinTemperature:   points[i].MinTemperature,
			WeatherCondition: mapWeatherCode(points[i].WeatherCode),
		}
	}

	return DailyWeatherResponse{
		Location: locationResp,
		Daily:    dailyResp,
		Units: DailyWeatherUnits{
			MaxTemperature: "°C",
			MinTemperature: "°C",
		},
	}, nil
}

func (s *WeatherService) GetByCity(city string) CityWeatherResponse {
	return CityWeatherResponse{
		City:        city,
		Temperature: 0,
		Condition:   "stub",
		Source:      "path",
	}
}

func mapWeatherCode(code int) string {
	mapper := map[int]string{
		0:  "Clear",
		1:  "Cloudy",
		2:  "Cloudy",
		3:  "Cloudy",
		45: "Fog",
		48: "Fog",
		51: "Drizzle",
		53: "Drizzle",
		55: "Drizzle",
		61: "Rain",
		63: "Rain",
		65: "Rain",
	}

	codeStr, ok := mapper[code]
	if !ok {
		return "Unknown"
	}

	return codeStr
}
