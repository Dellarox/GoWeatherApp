package main

import (
	"fmt"
	owm "github.com/briandowns/openweathermap"
	dotenv "github.com/joho/godotenv"
	"os"
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

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	err := dotenv.Load()
	checkError(err)

	owmApiKey := os.Getenv("OWM_API_KEY")

	w, err := owm.NewCurrent("C", "PL", owmApiKey) // fahrenheit (imperial) with Russian output
	checkError(err)

	err = w.CurrentByName("Nowa Sól,PL")
	checkError(err)

	weather := Weather{
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

	fmt.Printf("Miejsce: %s\n", weather.placeName)
	fmt.Printf("Temperatura: %.2f°C\n", weather.temperature)
	fmt.Printf("Opis pogody: %s\n", weather.description)
	fmt.Printf("Koordynaty (długość, szerokość): (%.3f, %.3f)\n",
		weather.coordinates.Longitude,
		weather.coordinates.Latitude)
	fmt.Printf("Wilgotność: %d%%\n", weather.humidity)
	fmt.Printf("Prędkość wiatru: %.2f m/s\n", weather.windSpeed)
}
