package weather

import (
	"WeatherApp/utils"
	"fmt"
	owm "github.com/briandowns/openweathermap"
	"time"
)

type Coordinates struct {
	Longitude float64
	Latitude  float64
}

type Weather struct {
	location    string
	temperature float64
	description string
	coordinates Coordinates
	humidity    int
	windSpeed   float64
}

func FetchWeather(apiKey, location string) Weather {
	w, err := owm.NewCurrent("C", "EN", apiKey)
	utils.CheckError(err)

	err = w.CurrentByName(location)
	utils.CheckError(err)

	return Weather{
		location:    w.Name,
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
		"Location: %s\nTemperature: %.2f°C\nWeather description: %s\nHumidity: %d%%\nWind speed: %.2f"+
			"m/s\nCoordinates: (%.3f, %.3f)",
		weather.location,
		weather.temperature,
		weather.description,
		weather.humidity,
		weather.windSpeed,
		weather.coordinates.Longitude,
		weather.coordinates.Latitude,
	)
}

func WriteLog(filename string, weather Weather) {
	writer, file := utils.CreateCSVWriter(filename)
	defer file.Close()
	defer writer.Flush()

	record := []string{
		weather.location,
		fmt.Sprintf("%.2f", weather.temperature),
		weather.description,
		fmt.Sprintf("%d", weather.humidity),
		fmt.Sprintf("%.2f", weather.windSpeed),
		fmt.Sprintf("(%.3f, %.3f)", weather.coordinates.Longitude, weather.coordinates.Latitude),
		time.Now().Format("2006-01-02_15-04-05"),
	}

	utils.WriteCSVRecord(writer, record)
}
