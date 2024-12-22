package weather

import (
	"WeatherApp/utils"
	"fmt"
	owm "github.com/briandowns/openweathermap"
)

type Coordinates struct {
	Longitude float64
	Latitude  float64
}

type Weather struct {
	placeName   string
	temperature float64
	description string
	coordinates Coordinates
	humidity    int
	windSpeed   float64
}

func FetchWeather(apiKey, location string) Weather {
	w, err := owm.NewCurrent("C", "PL", apiKey)
	utils.CheckError(err)

	err = w.CurrentByName(location)
	utils.CheckError(err)

	return Weather{
		placeName:   w.Name,
		temperature: w.Main.Temp,
		description: w.Weather[0].Description,
		coordinates: Coordinates{
			Longitude: w.GeoPos.Longitude,
			Latitude:  w.GeoPos.Latitude,
		},
		humidity:  w.Main.Humidity,
		windSpeed: w.Wind.Speed,
	}
}

func FormatWeather(weather Weather) string {
	return fmt.Sprintf(
		"Miejsce: %s\nTemperatura: %.2f°C\nOpis: %s\nWilgotność: %d%%\nWiatr: %.2f m/s\nKoordynaty: (%.3f, %.3f)",
		weather.placeName,
		weather.temperature,
		weather.description,
		weather.humidity,
		weather.windSpeed,
		weather.coordinates.Longitude,
		weather.coordinates.Latitude,
	)
}
