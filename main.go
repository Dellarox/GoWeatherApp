package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	owm "github.com/briandowns/openweathermap"
	dotenv "github.com/joho/godotenv"
	"os"
	"time"
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

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}

func setupClock() *widget.Label {
	clock := widget.NewLabel("")
	updateTime(clock)

	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

	return clock
}

func loadEnv() {
	err := dotenv.Load()
	checkError(err)
}

func fetchWeather(apiKey, location string) Weather {
	w, err := owm.NewCurrent("C", "PL", apiKey)
	checkError(err)

	err = w.CurrentByName(location)
	checkError(err)

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

func displayWeather(weather Weather, label *widget.Label) {
	formattedWeather := fmt.Sprintf(
		"Miejsce: %s\nTemperatura: %.2f°C\nOpis: %s\nWilgotność: %d%%\nWiatr: %.2f m/s\nKoordynaty: (%.3f, %.3f)",
		weather.placeName,
		weather.temperature,
		weather.description,
		weather.humidity,
		weather.windSpeed,
		weather.coordinates.Longitude,
		weather.coordinates.Latitude,
	)
	label.SetText(formattedWeather)
}

func main() {
	a := app.New()
	win := a.NewWindow("Clock & Weather")

	clock := setupClock()

	loadEnv()
	weatherLabel := widget.NewLabel("Loading Weather Information...")
	owmApiKey := os.Getenv("OWM_API_KEY")
	weather := fetchWeather(owmApiKey, "Złocieniec,PL")
	displayWeather(weather, weatherLabel)

	content := container.NewVBox(clock, weatherLabel)

	win.SetContent(content)
	win.ShowAndRun()
}
