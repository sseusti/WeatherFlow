package service

type WeatherService struct{}

type CurrentWeatherResponse struct {
	City        string
	Temperature int
	Condition   string
}

type CityWeatherResponse struct {
	City        string
	Temperature int
	Condition   string
	Source      string
}

func NewWeatherService() *WeatherService {
	return &WeatherService{}
}

func (s *WeatherService) GetCurrent(city string) CurrentWeatherResponse {
	return CurrentWeatherResponse{
		City:        city,
		Temperature: 0,
		Condition:   "stub",
	}
}

func (s *WeatherService) GetByCity(city string) CityWeatherResponse {
	return CityWeatherResponse{
		City:        city,
		Temperature: 0,
		Condition:   "stub",
		Source:      "ok",
	}
}
