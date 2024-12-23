package main

import (
	"WeatherApp/clock"
	"WeatherApp/utils"
	"WeatherApp/weather"
	"errors"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"os"
)

func main() {
	a := app.New()
	win := a.NewWindow("Clock & Weather")

	clockWidget := clock.SetupClock()

	utils.LoadEnv()
	weatherLabel := widget.NewLabel("Loading Weather Information...")
	owmApiKey := os.Getenv("OWM_API_KEY")

	locationEntry := widget.NewEntry()
	locationEntry.SetPlaceHolder("Enter the location")

	saveLocationButton := widget.NewButton("Save location", func() {
		if locationEntry.Text == "" {
			dialog.ShowError(errors.New("Please enter a location"), win)
		} else {
			weather.SetupWeather(owmApiKey, locationEntry.Text, weatherLabel)
		}
	})

	content := container.NewVBox(clockWidget, locationEntry, saveLocationButton, weatherLabel)

	win.SetContent(content)
	win.ShowAndRun()
}
