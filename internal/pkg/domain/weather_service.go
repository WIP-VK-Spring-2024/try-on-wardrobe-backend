package domain

//easyjson:json
type Weather struct {
	Temp float32 `json:"temp_c"`
}

//easyjson:json
type WeatherRequest struct {
	GeoPosition
	IP string
}

type WeatherService interface {
	CurrentWeather(WeatherRequest) (*Weather, error)
}
